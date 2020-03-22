package models

import "time"

type Metrics struct {
	UniqueVisitorsCount int64
	PublicationsCount   int64
	Timestamp           time.Time
}
