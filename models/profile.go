package models

import "time"

type Profile struct {
	ID           string
	UserID       string
	ProfileIsSet bool
	FirstName    string
	LastName     string
	Bio          string
	Followers    []string
	Following    []string
	Saved        []string
	Liked        []string
	ProfileImage string
	CreatedAt     time.Time
	Membership   *Membership
	Interests    []string
}
