package http

import (
	"edittapi/application/profile"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc profile.UseCase) {
	h := NewHandler(uc)

	profile := router.Group("/profile/")
	{
		profile.PATCH("", h.Update)
		profile.GET("", h.Get)
		//profile.GET(":id", h.GetByID)
		//profile.GET(":id/publications", h.GetPublications)
		profile.DELETE(":id", h.Delete)
	}
}
