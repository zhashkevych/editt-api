package http

import (
	"edittapi/pkg/feedback"
	"edittapi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	usecase feedback.UseCase
}

func newHandler(usecase feedback.UseCase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) createFeedback(c *gin.Context) {
	var inp models.Feedback
	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.usecase.CreateFeedback(c.Request.Context(), inp); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}