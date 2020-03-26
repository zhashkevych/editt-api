package delivery

import (
	"edittapi/pkg/admin"
	"edittapi/pkg/admin/delivery/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase admin.UseCase, authorizer *auth.Authorizer) {
	h := NewHandler(usecase, authorizer)

	router.POST("/sign-in", h.SignIn)

	router.GET("/metrics", authorizer.Middleware, h.GetMetrics)

	router.GET("/feedback", authorizer.Middleware, h.GetFeedbacks)

	publications := router.Group("/publications", authorizer.Middleware)
	{
		publications.GET("", h.GetPublications)
		publications.DELETE("/:id", h.RemovePublication)
	}
}
