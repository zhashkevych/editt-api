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
	Username string `json:"username" binding:"required" example:"editt"`
	Password string `json:"password" binding:"required" example:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

// SignIn godoc
// @Summary Sign In
// @Description Sign In
// @Tags admin
// @Accept json
// @Produce json
// @Param credentials body signInInput true "Sign In Input"
// @Success 200 {object} signInResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authorizer.GenerateToken(inp.Username, inp.Password)
	if err != nil {
		if err == auth.ErrInvalidAccessToken || err == auth.ErrInvalidCredentials {
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

// GetPublications godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer <token>"
// @Summary GetPublications
// @Description GetPublications
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} getPublicationsResponse
// @Failure 400
// @Failure 500
// @Router /admin/publications [get]
func (h *Handler) GetPublications(c *gin.Context) {
	ps, err := h.useCase.GetAllPublications(c.Request.Context())
	if err != nil {
		log.Errorf("error ocurred while getting publications from DB: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPublicationsResponse{ps})
}

// RemovePublication godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer <token>"
// @Summary RemovePublication
// @Description RemovePublication
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "Publication ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/publications/{id} [delete]
func (h *Handler) RemovePublication(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.RemovePublication(c.Request.Context(), id); err != nil {
		log.Errorf("error ocurred while removing publication from DB: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// GetMetrics godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer <token>"
// @Summary GetMetrics
// @Description GetMetrics
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} models.MetricsAggregated
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/metrics [get]
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

// GetFeedbacks godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer <token>"
// @Summary GetFeedbacks
// @Description GetFeedbacks
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} getFeedbacksResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/feedback [get]
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

