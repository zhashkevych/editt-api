package http

import (
	"edittapi/pkg/publication"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase) {
	h := NewHandler(usecase)

	publications := router.Group("/publications")
	{
		publications.POST("", h.Publish)
		publications.GET("", h.GetPublications)
		publications.GET("/:id", h.GetById)
		publications.POST("/:id/view", h.IncrementViews)
		publications.POST("/:id/reaction", h.IncrementReactions)
	}
}
