package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	accountsv1 "example.com/modmonolith/api/gen/accounts/v1"
	usersv1 "example.com/modmonolith/api/gen/users/v1"

	accountsapp "example.com/modmonolith/internal/modules/accounts/application/service"
	accountsgorm "example.com/modmonolith/internal/modules/accounts/infrastructure/gormrepo"
	accountsgrpc "example.com/modmonolith/internal/modules/accounts/interfaces/grpc"

	usersapp "example.com/modmonolith/internal/modules/users/application/service"
	usersgorm "example.com/modmonolith/internal/modules/users/infrastructure/gormrepo"
	usersgrpc "example.com/modmonolith/internal/modules/users/interfaces/grpc"
	userspublic "example.com/modmonolith/internal/modules/users/public"

	"example.com/modmonolith/internal/platform/config"
	"example.com/modmonolith/internal/platform/db"
	"example.com/modmonolith/internal/platform/grpcserver"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}
	isDev := cfg.Env != "prod"

	dbConn, err := db.New(ctx, cfg.PostgresDSN, isDev)
	if err != nil {
		log.Fatal(err)
	}

	// --- Users module
	usersRepo := usersgorm.New(dbConn.Gorm)
	if err := usersRepo.AutoMigrate(ctx); err != nil {
		log.Fatal(err)
	}
	usersHandlers := usersapp.NewHandlers(usersRepo)
	usersSvc := usersgrpc.New(usersHandlers)

	// Public contract adapter for other modules (in-process call)
	userReader := userspublic.NewUserReader(usersHandlers.Get)

	// --- Accounts module (depends on users via public contract only)
	accountsRepo := accountsgorm.New(dbConn.Gorm)
	if err := accountsRepo.AutoMigrate(ctx); err != nil {
		log.Fatal(err)
	}
	accountsHandlers := accountsapp.NewHandlers(accountsRepo, userReader)
	accountsSvc := accountsgrpc.New(accountsHandlers)

	// --- gRPC server
	srv, err := grpcserver.New(cfg.GRPCAddr,
		grpcserver.UnaryChain(
			grpcserver.UnaryRecovery(),
			grpcserver.UnaryLogging(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	usersv1.RegisterUsersServiceServer(srv.GRPC(), usersSvc)
	accountsv1.RegisterAccountsServiceServer(srv.GRPC(), accountsSvc)

	go func() {
		log.Printf("gRPC listening on %s", cfg.GRPCAddr)
		if err := srv.Serve(); err != nil {
			log.Printf("grpc serve: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	log.Printf("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Stop(shutdownCtx); err != nil {
		log.Printf("grpc stop: %v", err)
		os.Exit(1)
	}
}
