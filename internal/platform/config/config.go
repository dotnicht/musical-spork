package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env string

	GRPCAddr string
	HTTPAddr string

	PostgresDSN string
}

func FromEnv() (Config, error) {
	env := getenv("APP_ENV", "dev")
	grpcAddr := getenv("GRPC_ADDR", ":50051")
	httpAddr := getenv("HTTP_ADDR", ":8080")

	dsn := getenv("POSTGRES_DSN",
		"host=localhost user=app password=app dbname=app port=5432 sslmode=disable TimeZone=UTC",
	)
	if dsn == "" {
		return Config{}, fmt.Errorf("POSTGRES_DSN is required")
	}

	return Config{
		Env:         env,
		GRPCAddr:    grpcAddr,
		HTTPAddr:    httpAddr,
		PostgresDSN: dsn,
	}, nil
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
