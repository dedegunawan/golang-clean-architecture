package server

import (
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

func New(lg *logger.Logger) *Engine {
	g := gin.New()
	return &Engine{g}
}
