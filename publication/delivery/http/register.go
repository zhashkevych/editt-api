package http

import (
	"edittapi/publication"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase) {
	h := NewHandler(usecase)

	bookmarks := router.Group("/publications")
	{
		bookmarks.POST("", h.Publish)
		bookmarks.GET("", h.GetPublications)
		bookmarks.GET("/:id", h.GetById)
	}
}
