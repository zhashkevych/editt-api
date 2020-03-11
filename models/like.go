package models

type Like struct {
	ID     string
	Author *Profile
	Count  int32
}
