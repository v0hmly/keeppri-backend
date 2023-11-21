package main

import (
	fLog "log"
	"log/slog"
	"os"
	"strconv"

	"github.com/v0hmly/keeppri-backend/internal/config"
	log "github.com/v0hmly/keeppri-backend/internal/lib/logger"
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

	return nil
}
