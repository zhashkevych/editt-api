package models

type Feedback struct {
	Score    int32 `json:"score" example:"10"`
	Features []int32 `json:"features" enums:"1,2"`
}
