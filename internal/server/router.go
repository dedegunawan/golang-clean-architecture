package server

import (
	"github.com/gin-gonic/gin"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
	"github.com/dedegunawan/golang-clean-architecture/internal/server/middleware"
)

type Engine struct {
	*gin.Engine
}

func New(lg *logger.Logger) *Engine {
	g := gin.New()
	g.Use(middleware.RequestID())
	g.Use(middleware.ZapLogger(lg))
	g.Use(middleware.Recovery(lg))
	return &Engine{g}
}
