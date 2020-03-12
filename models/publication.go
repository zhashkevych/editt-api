package models

import "time"

type Publication struct {
	ID          string
	Author      string
	Tags        []string
	Body        string
	ImageLink   string
	Views       int32
	Claps       int32
	ReadingTime int32
	PublishedAt time.Time
}
