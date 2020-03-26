package http

import (
	"edittapi/pkg/feedback"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPHandlers(router *gin.RouterGroup, usecase feedback.UseCase) {
	h := newHandler(usecase)

	router.POST("/feedback", h.createFeedback)
}