package models

import "time"

type MembershipType int

const (
	MembershipTypeFree MembershipType = iota
	MembershipTypePaid
)

type Membership struct {
	ID         string
	Type       MembershipType
	isExpiried bool
	StartedAt  time.Time
	ExpiresAt  time.Time
}
