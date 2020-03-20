package http

import (
	"edittapi/pkg/publication"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase) {
	h := NewHandler(usecase)

	bookmarks := router.Group("/publications")
	{
		bookmarks.POST("", h.Publish)
		bookmarks.GET("", h.GetPublications)
		bookmarks.GET("/:id", h.GetById)
		bookmarks.POST("/:id/view", h.IncrementViews)
		bookmarks.POST("/:id/reaction", h.IncrementReactions)
	}
}
