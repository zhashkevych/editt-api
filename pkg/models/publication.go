package models

import "time"

type Publication struct {
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Tags        []string  `json:"tags"`
	Body        string    `json:"body"`
	ImageLink   string    `json:"imageLink"`
	Reactions   int32     `json:"reactions"`
	Views       int32     `json:"views"`
	ReadingTime int32     `json:"readingTime"`
	PublishedAt time.Time `json:"publishedAt"`
}
