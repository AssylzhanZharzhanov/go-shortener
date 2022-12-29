package http

import (
	"net/http"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/helpers"
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"

	"github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
)

// Handler - Impl.
type Handler struct {
	service domain.ShortenerService
}

// NewHandler - returns a new handler
func NewHandler(service domain.ShortenerService) *Handler {
	return &Handler{service: service}
}

// Create - Impl.
// @Summary Create shorten url
// @Tags news
// @Description create shorten url
// @ID create-shorten-url
// @Accept json
// @Produce json
// @Param input body domain.CreateShortenURLDTO true "news body"
// @Success 200 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/shorten [post]
func (h *Handler) Create(c *gin.Context) {
	var input domain.CreateShortenURLDTO
	if err := c.BindJSON(&input); err != nil {
		return
	}

	var newLink *domain.Link
	err := dto.Map(&newLink, input)
	if err != nil {
		helpers.NewResponse(c, http.StatusInternalServerError, helpers.ResponseStatusError, err)
		return
	}

	response, err := h.service.CreateShortenURL(c.Request.Context(), newLink)
	if err != nil {
		helpers.NewResponse(c, http.StatusInternalServerError, helpers.ResponseStatusError, err)
		return
	}

	helpers.NewResponse(c, http.StatusOK, helpers.ResponseStatusSuccess, response)
}

// Get - Impl.
// @Summary Get url by id
// @Tags news
// @Description Get url by id
// @ID get-url-by-id
// @Accept json
// @Produce json
// @Param shorten_url path string true "shorten url"
// @Success 200 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/shorten/{shorten_url} [get]
func (h *Handler) Get(c *gin.Context) {
	var (
		shortenUrl = c.Param("id")
	)

	response, err := h.service.Get(c.Request.Context(), shortenUrl)
	if err != nil {
		helpers.NewResponse(c, http.StatusInternalServerError, helpers.ResponseStatusError, err)
		return
	}

	c.Redirect(http.StatusFound, response.OriginalURL)
}
