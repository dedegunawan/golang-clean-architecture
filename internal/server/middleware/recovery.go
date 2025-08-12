package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dedegunawan/golang-clean-architecture/pkg/logger"
	"github.com/dedegunawan/golang-clean-architecture/pkg/response"
)

func Recovery(lg *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				lg.Errorw("panic", "error", rec)
				response.Error(c, http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
