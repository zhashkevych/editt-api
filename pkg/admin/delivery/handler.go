package delivery

import (
	"edittapi/pkg/admin"
	"edittapi/pkg/admin/delivery/auth"
	"edittapi/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

		log.Errorf("error ocurred while generating JWT Token: %s", err.Error())
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
		log.Errorf("error ocurred while getting publications from DB: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPublicationsResponse{ps})
}

func (h *Handler) RemovePublication(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.RemovePublication(c.Request.Context(), id); err != nil {
		log.Errorf("error ocurred while removing publication from DB: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetMetrics(c *gin.Context) {
	metricsData, err := h.useCase.GetMetrics(c.Request.Context())
	if err != nil {
		log.Errorf("error ocurred while getting metrics: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, metricsData)
}

type getFeedbacksResponse struct {
	Feedbacks []*models.Feedback `json:"feedbacks"`
}

func (h *Handler) GetFeedbacks(c *gin.Context) {
	feedbacks, err := h.useCase.GetFeedbacks(c.Request.Context())
	if err != nil {
		log.Errorf("error ocurred while getting feedbacks: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getFeedbacksResponse{
		feedbacks,
	})
}

