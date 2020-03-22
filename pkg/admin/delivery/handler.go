package delivery

import (
	"edittapi/pkg/admin"
	"edittapi/pkg/admin/delivery/auth"
	"edittapi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	authorizer *auth.Authorizer
	useCase    admin.UseCase
}

func NewHandler(useCase admin.UseCase, authorizer *auth.Authorizer) *Handler {
	return &Handler{
		useCase:    useCase,
		authorizer: authorizer,
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authorizer.GenerateToken(inp.Username, inp.Password)
	if err != nil {
		if err == auth.ErrInvalidAccessToken {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &signInResponse{
		Token: token,
	})
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
	Last24Hours       *models.Metrics `json:"last24"`
	LastHour          *models.Metrics `json:"lastHour"`
	PublicationsCount int64           `json:"publicationsCount"`
}

func (h *Handler) GetMetrics(c *gin.Context) {
	ms, err := h.useCase.GetMetrics(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getMetricsResponse{
		LastHour:          ms.LastHour,
		Last24Hours:       ms.Last24Hours,
		PublicationsCount: ms.PublicationsCount,
	})
}
