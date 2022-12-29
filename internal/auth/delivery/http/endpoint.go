package http

import (
	"github.com/AssylzhanZharzhanov/go-shortener/internal/auth/domain"

	"github.com/gin-gonic/gin"
)

// RegisterEndpoints - registers new endpoints
func RegisterEndpoints(router *gin.RouterGroup, service domain.AuthService) {
	h := NewHandler(service)

	users := router.Group("/auth")
	{
		users.POST("/", h.Register)
		users.POST("/:id", h.Authenticate)
	}
}
