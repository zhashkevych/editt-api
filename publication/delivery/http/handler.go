package http

import (
	"edittapi/models"
	"edittapi/publication"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	useCase publication.UseCase
}

func NewHandler(useCase publication.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type publishInput struct {
	Author    string   `json:"author" binding:"required,min=3,max=25"`
	Tags      []string `json:"tags" binding:"required,min=1,max=3"`
	Body      string   `json:"body" binding:"required"`
	ImageLink string   `json:"imageLink" binding:"required"`
}

func toPublicationModel(inp *publishInput) models.Publication {
	return models.Publication{
		Author:    inp.Author,
		Tags:      inp.Tags,
		Body:      inp.Body,
		ImageLink: inp.ImageLink,
	}
}

func (h *Handler) Publish(c *gin.Context) {
	inp := new(publishInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	p := toPublicationModel(inp)

	if err := h.useCase.Publish(c.Request.Context(), p); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type publicationResponse struct {
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags"`
	Body        string    `json:"body"`
	ImageLink   string    `json:"imageLink"`
	Claps       int32     `json:"claps"`
	ReadingTime int32     `json:"readingTime"`
	PublishedAt time.Time `json:"publishedAt"`
}

type getPublicationsResponse struct {
	Publications []*publicationResponse `json:"publications"`
}

func (h *Handler) GetPopular(c *gin.Context) {
	limit := c.DefaultQuery("limit", "3")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Errorf("failed to parse 'limit' query parameter: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ps, err := h.useCase.GetPopularPublications(c.Request.Context(), int64(limitInt))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPublicationsResponse{
		Publications: toPublications(ps),
	})
}

func (h *Handler) GetLatest(c *gin.Context) {
	limit := c.DefaultQuery("limit", "3")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Errorf("failed to parse 'limit' query parameter: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ps, err := h.useCase.GetLatestPublications(c.Request.Context(), int64(limitInt))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getPublicationsResponse{
		Publications: toPublications(ps),
	})
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	p, err := h.useCase.GetById(c.Request.Context(), id)
	if err != nil {
		log.Errorf("failed to parse 'limit' query parameter: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, toPublication(p))
}

func toPublication(p *models.Publication) *publicationResponse {
	return &publicationResponse{
		ID:          p.ID,
		Author:      p.Author,
		Tags:        p.Tags,
		Body:        p.Body,
		ImageLink:   p.ImageLink,
		Claps:       p.Claps,
		ReadingTime: p.ReadingTime,
		PublishedAt: p.PublishedAt,
	}
}

func toPublications(ps []*models.Publication) []*publicationResponse {
	out := make([]*publicationResponse, len(ps))

	for i := range ps {
		out[i] = toPublication(ps[i])
	}

	return out
}
