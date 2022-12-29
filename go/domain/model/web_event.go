package model

import "time"

type WebEvent struct {
	ID        int
	AccountID int
	Account   Account
	OccuredAt time.Time
	Channel   string
}
