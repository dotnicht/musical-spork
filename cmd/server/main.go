package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	usersv1 "example.com/modmonolith/api/gen/users/v1"
	"example.com/modmonolith/internal/modules/users/application/service"
	"example.com/modmonolith/internal/modules/users/infrastructure/gormrepo"
	usersgrpc "example.com/modmonolith/internal/modules/users/interfaces/grpc"
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

	userRepo := gormrepo.New(dbConn.Gorm)
	if err := userRepo.AutoMigrate(ctx); err != nil {
		log.Fatal(err)
	}

	usersHandlers := service.NewHandlers(userRepo)
	usersSvc := usersgrpc.New(usersHandlers)

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
