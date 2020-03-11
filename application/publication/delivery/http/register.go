package http

import (
	"edittapi/application/publication"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, puc publication.PublicationUseCase,
	cuc publication.CommentUseCase, luc publication.LikeUseCase) {
	h := NewHandler(puc, cuc, luc)

	bookmarks := router.Group("/publication/")
	{
		bookmarks.POST("", h.Publish)
		bookmarks.GET(":id", h.Get)
		bookmarks.DELETE("", h.Delete)
		bookmarks.PATCH(":id", h.Update)
	}
}
