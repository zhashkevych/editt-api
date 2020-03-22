package http

import (
	"edittapi/pkg/publication"
	"edittapi/pkg/publication/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase publication.UseCase) {
	h := NewHandler(usecase)

	viewer := middleware.NewViewer(usecase)

	publications := router.Group("/publications")
	{
		publications.POST("", h.Publish)
		publications.GET("", h.GetPublications)
		publications.GET("/:id", viewer.Middleware, h.GetById)
		publications.POST("/:id/reaction", h.IncrementReactions)
	}
}
