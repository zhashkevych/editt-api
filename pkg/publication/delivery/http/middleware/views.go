package middleware

import (
	"edittapi/pkg/publication"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type Viewer struct {
	useCase publication.UseCase
	IPs     map[string]interface{}
	sync.Mutex
}

func NewViewer(useCase publication.UseCase) *Viewer {
	return &Viewer{
		useCase: useCase,
		IPs:     make(map[string]interface{}),
	}
}

func (v *Viewer) Middleware(c *gin.Context) {
	id := c.Param("id")

	// Increment Views by IP Address
	v.Lock()
	defer v.Unlock()
	if _, ex := v.IPs[id]; ex {
		return
	}

	if err := v.useCase.IncrementViews(c.Request.Context(), id); err != nil {
		log.Errorf("error occured while increasing views: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	v.IPs[id] = nil
}
