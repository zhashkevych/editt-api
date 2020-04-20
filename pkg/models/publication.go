package models

import "time"

type Publication struct {
	ID          string    `json:"id" example:"507f1f77bcf86cd799439011"`
	Author      string    `json:"author" example:"Вася"`
	Title       string    `json:"title" example:"Про личные финансы"`
	Tags        []string  `json:"tags" example:"финансы,бюджет"`
	Body        string    `json:"body" example:"Очень крутая публикация"`
	ImageLink   string    `json:"imageLink" example:"https://images.unsplash.com/photo-1571997804104-011c8c1d19b6?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1650&q=80"`
	Reactions   int32     `json:"reactions" example:"35"`
	Views       int32     `json:"views" example:"586"`
	ReadingTime int32     `json:"readingTime" example:"5"`
	PublishedAt time.Time `json:"publishedAt"`
}
