package delivery

import (
	"edittapi/pkg/admin"
	"edittapi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase admin.UseCase
}

func NewHandler(useCase admin.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type getPublicationsResponse struct {
	Publications []*models.Publication `json:"publications"`
}

func (h *Handler) GetPublications(c *gin.Context) {
	ps, err := h.useCase.GetAllPublications(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPublicationsResponse{ps})
}

func (h *Handler) RemovePublication(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.RemovePublication(c.Request.Context(), id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type getMetricsResponse struct {
	Metrics []*models.Metrics `json:"metrics"`
}

func (h *Handler) GetMetrics(c *gin.Context) {
	ms, err := h.useCase.GetMetrics(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getMetricsResponse{ms})
}
