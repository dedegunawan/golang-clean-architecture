package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data any) {
	c.JSON(200, APIResponse{Success: true, Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(201, APIResponse{Success: true, Data: data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, APIResponse{Success: false, Message: msg})
}
