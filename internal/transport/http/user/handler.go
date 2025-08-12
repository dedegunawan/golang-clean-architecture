package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/dedegunawan/golang-clean-architecture/internal/domain/user"
	"github.com/dedegunawan/golang-clean-architecture/pkg/response"
)

type registerReq struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type avatarReq struct {
	Avatar string `json:"avatar" binding:"required,url"`
}

type activeReq struct {
	IsActive bool `json:"is_active"`
}

func RegisterRoutes(r *gin.Engine, svc domain.Service) {
	g := r.Group("/v1/users")

	g.POST("", func(c *gin.Context) {
		var req registerReq
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error()); return
		}
		u, err := svc.Register(req.Name, req.Email, req.Password)
		if err != nil { response.Error(c, 400, err.Error()); return }
		response.Created(c, u)
	})

	g.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		u, err := svc.Get(id)
		if err != nil { response.Error(c, 404, "not found"); return }
		response.OK(c, u)
	})

	g.GET("", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
		items, total, err := svc.List(page, size)
		if err != nil { response.Error(c, 400, err.Error()); return }
		response.OK(c, gin.H{"items": items, "total": total, "page": page, "size": size})
	})

	g.PUT("/:id/avatar", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var req avatarReq
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error()); return
		}
		if err := svc.UpdateAvatar(id, req.Avatar); err != nil {
			response.Error(c, 400, err.Error()); return
		}
		response.OK(c, gin.H{"id": id, "avatar": req.Avatar})
	})

	g.PUT("/:id/active", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var req activeReq
		if err := c.ShouldBindJSON(&req); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error()); return
		}
		if err := svc.SetActive(id, req.IsActive); err != nil {
			response.Error(c, 400, err.Error()); return
		}
		response.OK(c, gin.H{"id": id, "is_active": req.IsActive})
	})
}
