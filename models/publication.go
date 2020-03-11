package models

import "time"

type Publication struct {
	ID          string
	Author 		*Profile
	S3Link      string
	ViewsCount  int32
	Likes       []*Like
	Comments    []*Comment
	PublishedAt time.Time
	Tags        []string
}
