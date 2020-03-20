package delivery

import (
	"edittapi/pkg/admin"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase admin.UseCase) {
	h := NewHandler(usecase)

	router.GET("/metrics", h.GetMetrics)

	publications := router.Group("/publications")
	{
		publications.GET("", h.GetPublications)
		publications.DELETE("/:id", h.RemovePublication)
	}
}

