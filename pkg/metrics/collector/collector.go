package collector

import (
	"context"
	"edittapi/pkg/metrics"
	"edittapi/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type MetricsCollector struct {
	sync.Mutex
	IPs     map[string]interface{}
	useCase metrics.UseCase
}

func NewMetricsCollector(usecase metrics.UseCase) *MetricsCollector {
	return &MetricsCollector{
		IPs:     make(map[string]interface{}),
		useCase: usecase,
	}
}

func (mc *MetricsCollector) Middleware(c *gin.Context) {
	mc.Lock()
	mc.IPs[c.ClientIP()] = nil
	mc.Unlock()
}

func (mc *MetricsCollector) Flush(ctx context.Context) {
	ms := models.Metrics{
		UniqueVisitorsCount: int64(len(mc.IPs)),
		Timestamp:           time.Now(),
	}

	if err := mc.useCase.SetMetrics(ctx, ms); err != nil {
		log.Errorf("failed to write metrics: %s", err.Error())
		return
	}

	mc.Lock()
	mc.IPs = make(map[string]interface{})
	mc.Unlock()
}
