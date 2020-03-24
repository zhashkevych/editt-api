package http

import (
	"edittapi/pkg/publication"
	"edittapi/pkg/publication/delivery/http/middleware"
	"edittapi/sidecar/filestorage"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase, fileStorage *filestorage.FileStorage) {
	h := NewHandler(usecase, fileStorage)

	viewer := middleware.NewViewer(usecase)

	router.POST("/upload", h.Upload)

	publications := router.Group("/publications")
	{
		publications.POST("", h.Publish)
		publications.GET("", h.GetPublications)
		publications.GET("/:id", viewer.Middleware, h.GetById)
		publications.POST("/:id/reaction", h.IncrementReactions)
	}
}
