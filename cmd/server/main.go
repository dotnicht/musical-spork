package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	accountsv1 "example.com/modmonolith/api/gen/accounts/v1"
	usersv1 "example.com/modmonolith/api/gen/users/v1"

	accountsapp "example.com/modmonolith/internal/modules/accounts/application/service"
	accountsgorm "example.com/modmonolith/internal/modules/accounts/infrastructure/gormrepo"
	accountsgrpc "example.com/modmonolith/internal/modules/accounts/interfaces/grpc"

	usersapp "example.com/modmonolith/internal/modules/users/application/service"
	usersgorm "example.com/modmonolith/internal/modules/users/infrastructure/gormrepo"
	usersgrpc "example.com/modmonolith/internal/modules/users/interfaces/grpc"
	usershttp "example.com/modmonolith/internal/modules/users/interfaces/http"
	userspublic "example.com/modmonolith/internal/modules/users/public"

	"example.com/modmonolith/internal/platform/config"
	"example.com/modmonolith/internal/platform/db"
	"example.com/modmonolith/internal/platform/grpcserver"
	"example.com/modmonolith/internal/platform/httpserver"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.FromEnv()
	if err != nil { log.Fatal(err) }
	isDev := cfg.Env != "prod"

	dbConn, err := db.New(ctx, cfg.PostgresDSN, isDev)
	if err != nil { log.Fatal(err) }

	// --- Users module
	usersRepo := usersgorm.New(dbConn.Gorm)
	if err := usersRepo.AutoMigrate(ctx); err != nil { log.Fatal(err) }
	usersHandlers := usersapp.NewHandlers(usersRepo)
	usersGRPCSvc := usersgrpc.New(usersHandlers)
	usersHTTP := usershttp.New(usersHandlers)

	// Public contract for other modules (in-process)
	userReader := userspublic.NewUserReader(usersHandlers.Get)

	// --- Accounts module
	accountsRepo := accountsgorm.New(dbConn.Gorm)
	if err := accountsRepo.AutoMigrate(ctx); err != nil { log.Fatal(err) }
	accountsHandlers := accountsapp.NewHandlers(accountsRepo, userReader)
	accountsGRPCSvc := accountsgrpc.New(accountsHandlers)

	// --- gRPC server
	grpcSrv, err := grpcserver.New(cfg.GRPCAddr,
		grpcserver.UnaryChain(
			grpcserver.UnaryRecovery(),
			grpcserver.UnaryLogging(),
		),
	)
	if err != nil { log.Fatal(err) }
	usersv1.RegisterUsersServiceServer(grpcSrv.GRPC(), usersGRPCSvc)
	accountsv1.RegisterAccountsServiceServer(grpcSrv.GRPC(), accountsGRPCSvc)

	// --- HTTP server
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	r.Mount("/", usersHTTP.Routes())
	httpSrv := httpserver.New(cfg.HTTPAddr, r)

	go func() {
		log.Printf("gRPC listening on %s", cfg.GRPCAddr)
		if err := grpcSrv.Serve(); err != nil {
			log.Printf("grpc serve: %v", err)
			stop()
		}
	}()

	go func() {
		log.Printf("HTTP listening on %s", cfg.HTTPAddr)
		if err := httpSrv.Serve(); err != nil {
			log.Printf("http serve: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Printf("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = httpSrv.Stop(shutdownCtx)
	_ = grpcSrv.Stop(shutdownCtx)

	log.Printf("bye")
	os.Exit(0)
}
