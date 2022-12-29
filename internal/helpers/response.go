package helpers

import "github.com/gin-gonic/gin"

// Response statuses.
const (
	ResponseStatusSuccess = ResponseStatus("success")
	ResponseStatusError   = ResponseStatus("error")
)

// ResponseStatus - represents the response status.
type ResponseStatus string

// Response - represents the default response
type Response struct {
	Body struct {
		Status ResponseStatus `json:"status"`
		Data   interface{}    `json:"data"`
	} `json:"body"`
}

// NewResponse - creates a response and send it to client
func NewResponse(c *gin.Context, statusCode int, status ResponseStatus, data interface{}) {
	c.JSON(statusCode, Response{
		Body: struct {
			Status ResponseStatus `json:"status"`
			Data   interface{}    `json:"data"`
		}{
			Status: status,
			Data:   data,
		},
	})
}
