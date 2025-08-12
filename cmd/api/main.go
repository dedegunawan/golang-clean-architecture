package main

import (
	"github.com/dedegunawan/golang-clean-architecture/config"
	"github.com/dedegunawan/golang-clean-architecture/internal/app"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
	"os"
)

func main() {
	// load .env (no-op jika tidak ada)
	_ = config.LoadDotEnv()
	cfg := config.FromEnv()

	// Pastikan folder log ada
	os.MkdirAll("logs", 0755)
	
	// init logger
	lg := logger.New(cfg.LogLevel, "logs/app.log")
	defer lg.Sync()

	// bootstrap app (DB, router, wiring modul)
	a, err := app.New(cfg, lg)
	if err != nil {
		lg.Fatalw("bootstrap failed", "error", err)
	}

	// start server the Gin way
	lg.Infow("server starting (gin)", "addr", cfg.HTTPAddr)
	if err := a.Router.Engine.Run(cfg.HTTPAddr); err != nil {
		lg.Fatalw("listen failed", "error", err)
	}
}
