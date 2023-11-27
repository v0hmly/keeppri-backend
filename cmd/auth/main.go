package main

import (
	"fmt"
	fLog "log"
	"log/slog"
	"os"
	"strconv"

	"github.com/v0hmly/keeppri-backend/internal/config"
	"github.com/v0hmly/keeppri-backend/internal/grpc"
	"github.com/v0hmly/keeppri-backend/internal/lib/hash"
	log "github.com/v0hmly/keeppri-backend/internal/lib/logger"
	"github.com/v0hmly/keeppri-backend/internal/lib/token"
	"github.com/v0hmly/keeppri-backend/internal/repository"
	"github.com/v0hmly/keeppri-backend/internal/services"
)

func main() {
	if err := run(); err != nil {
		fLog.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	err, cfg := config.MustLoad()
	if err != nil {
		return err
	}

	logger := log.SetupLogger(cfg.Env)

	logger.Info(
		"starting Keeppri auth service",
		slog.String("production: ", strconv.FormatBool(cfg.Env == "prod")),
		slog.String("version: ", cfg.Version))

	logger.Debug("debug messages are enabled")

	repo, err := repository.New(cfg)
	if err != nil {
		return err
	}

	newServices := services.NewServices(services.Deps{
		Logger:       logger,
		Repos:        repo,
		Hash:         hash.NewBcryptPasswordHasher(),
		TokenManager: token.NewSessionTokenGenerator(cfg.Token.SessionTokenSize),
	})

	grpcHandler := grpc.NewGrpcHandler(newServices)

	if err = grpcHandler.Run(":" + cfg.GRPC.Port); err != nil {
		return fmt.Errorf("grpc server failed: %w", err)
	}

	return nil
}
