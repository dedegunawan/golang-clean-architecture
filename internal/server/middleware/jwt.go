package middleware

import (
	"github.com/dedegunawan/golang-clean-architecture/pkg/jwtmanager"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ContextKey untuk set/get dari Gin Context
const (
	ContextUserID = "user_id"
	ContextEmail  = "user_email"
	ContextRole   = "user_role"
)

func AuthJWT(mgr *jwtmanager.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "missing bearer token"})
			return
		}
		token := strings.TrimPrefix(h, "Bearer ")
		claims, err := mgr.Validate(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "invalid token"})
			return
		}
		// simpan user id ke context
		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextEmail, claims.Email)
		c.Set(ContextRole, claims.Role)
		c.Next()
	}
}
