package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dedegunawan/golang-clean-architecture/config"
	"github.com/dedegunawan/golang-clean-architecture/internal/app"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
)

func main() {
	// load env
	if err := config.LoadDotEnv(); err != nil {
		// ignore when not present
	}

	cfg := config.FromEnv()

	// logger
	lg := logger.New(cfg.LogLevel)
	defer lg.Sync()

	// init app (db, router, wiring)
	a, err := app.New(cfg, lg)
	if err != nil {
		lg.Fatalw("bootstrap failed", "error", err)
	}

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: a.Router, // gin.Engine
	}

	go func() {
		lg.Infow("server starting", "addr", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lg.Fatalw("listen failed", "error", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		lg.Errorw("shutdown error", "error", err)
	}
	lg.Info("server stopped")
}
