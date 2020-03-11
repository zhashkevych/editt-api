package http

import (
	"edittapi/application/profile"
	"github.com/gin-gonic/gin"
	"edittapi/auth"
	"edittapi/models"
	"net/http"
	"time"
)

type Bookmark struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Handler struct {
	useCase profile.UseCase
}

func NewHandler(useCase profile.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Update(c *gin.Context) {
	inp := profile.ProfileInput{}
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.UpdateProfile(c.Request.Context(), user, inp); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type profileResponse struct {
	ProfileIsSet bool               `json:"isSet"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	Bio          string             `json:"bio"`
	Followers    []string           `json:"followers"`
	Following    []string           `json:"following"`
	Saved        []string           `json:"saved"`
	Liked        []string           `json:"liked"`
	ProfileImage string             `json:"profileImage"`
	CreatedAt    time.Time          `json:"createdAt"`
	Membership   *models.Membership `json:"membership"`
	Interests    []string           `json:"interests"`
}

func toProfileResponse(p *models.Profile) *profileResponse {
	return &profileResponse{
		ProfileIsSet: p.ProfileIsSet,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Bio:          p.Bio,
		Followers:    p.Followers,
		Following:    p.Following,
		Saved:        p.Saved,
		Liked:        p.Liked,
		ProfileImage: p.ProfileImage,
		CreatedAt:    p.CreatedAt,
		Membership:   &models.Membership{},
		Interests:    p.Interests,
	}
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	c.JSON(http.StatusOK, toProfileResponse(user.Profile))
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteProfile(c.Request.Context(), user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
