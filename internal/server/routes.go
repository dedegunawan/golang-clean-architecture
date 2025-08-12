package server

import (
	"github.com/dedegunawan/golang-clean-architecture/internal/transport/http/auth"
	"github.com/dedegunawan/golang-clean-architecture/internal/transport/http/user"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth *auth.AuthHandler
	User *user.UserHandler
	// nanti tambah lagi misalnya Product, Order, dsb.
}
type Middleware struct {
	AuthJWT   gin.HandlerFunc
	CORS      gin.HandlerFunc
	RequestID gin.HandlerFunc
	Logger    gin.HandlerFunc
	Recovery  gin.HandlerFunc
	Rate      gin.HandlerFunc
	// nanti tambah lagi misalnya Product, Order, dsb.
}

func RegisterRoutes(r *gin.Engine, h *Handlers, m *Middleware) {

	r.Use(m.CORS, m.RequestID, m.Logger, m.Recovery)

	api := r.Group("/api/v1")

	protected := api.Group("")
	protected.Use(m.AuthJWT) // tinggal pakai

	// public route
	h.Auth.RegisterRoute(api)

	// protected route
	h.User.RegisterRoute(protected)

	api.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"success": true, "status": "ok"}) })

}
