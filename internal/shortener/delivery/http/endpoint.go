package http

import (
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
	"github.com/gin-gonic/gin"
)

// RegisterEndpoints - registers new endpoints
func RegisterEndpoints(router *gin.RouterGroup, service domain.ShortenerService) {
	h := NewHandler(service)

	shortener := router.Group("/shorten")
	{
		shortener.POST("/", h.Create)
		shortener.GET("/:shorten_url", h.Get)
		//shortener.DELETE("/:id", h.Delete)
	}
}
