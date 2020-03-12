package http

import (
	"edittapi/publication"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase) {
	h := NewHandler(usecase)

	bookmarks := router.Group("/publication")
	{
		bookmarks.POST("", h.Publish)
		bookmarks.GET("/popular", h.GetPopular)
		bookmarks.GET("/latest", h.GetLatest)
	}
}
