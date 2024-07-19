package http

import (
	"net/http"

	"github.com/crseat/example-data-pipeline/internal/application"
	"github.com/crseat/example-data-pipeline/internal/domain"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	service *application.PostService
}

func NewHandler(service *application.PostService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/submit", h.handleSubmit)
}

func (h *Handler) handleSubmit(c echo.Context) error {
	var postData domain.PostData

	// Bind and validate the JSON data
	if err := c.Bind(&postData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	// Validate the struct
	if err := c.Validate(&postData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Process the post data
	if err := h.service.ProcessPostData(postData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to process data"})
	}

	// No content response if data is valid
	return c.NoContent(http.StatusNoContent)
}
