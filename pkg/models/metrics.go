package models

import "time"

type Metrics struct {
	UniqueVisitorsCount int32
	Timestamp           time.Time
}
