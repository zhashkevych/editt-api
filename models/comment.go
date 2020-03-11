package models

import "time"

type Comment struct {
	ID          string
	Author      *Profile
	Text        string
	PublishedAt time.Time
}
