package http

import (
	"edittapi/application/feed"
	"github.com/gin-gonic/gin"
	"edittapi/bookmark"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc feed.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/feed")
	{
		bookmarks.GET("", h.Get)
	}
}
