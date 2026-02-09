package main

import (
	"log/slog"

	"github.com/koha90/tui_bots_manager_v1/internal/config"
	"github.com/koha90/tui_bots_manager_v1/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.SetupLogger(cfg.App.LogLevel)
	slog.SetDefault(logger)

	slog.Info("application: " + cfg.Name + " version: " + cfg.Version + " is started")
	slog.Debug("debug mod is anabled")
}
