package models

type Feedback struct {
	Score    int32 `json:"score"`
	Features []int32 `json:"features"`
}
