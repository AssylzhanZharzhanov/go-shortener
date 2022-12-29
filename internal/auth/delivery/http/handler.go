package http

import (
	"github.com/AssylzhanZharzhanov/go-shortener/internal/auth/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler - Impl.
type Handler struct {
	service domain.AuthService
}

// NewHandler - returns hundler
func NewHandler(service domain.AuthService) *Handler {
	return &Handler{service: service}
}

// Register - Impl.
func (h *Handler) Register(c *gin.Context) {

	//var input *domain.CreateRoomRequestDTO
	//if err := c.BindJSON(&input); err != nil {
	//	return
	//}
	//
	//room, err := h.service.CreateRoom(c.Request.Context(), input)
	//if err != nil {
	//	return
	//}

	c.JSON(http.StatusCreated, "room")
}

// Authenticate - Impl.
func (h *Handler) Authenticate(c *gin.Context) {

	//var input *domain.CreateRoomRequestDTO
	//if err := c.BindJSON(&input); err != nil {
	//	return
	//}
	//
	//room, err := h.service.CreateRoom(c.Request.Context(), input)
	//if err != nil {
	//	return
	//}

	c.JSON(http.StatusCreated, "room")
}
