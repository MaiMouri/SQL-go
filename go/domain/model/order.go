package model

import "time"

type Order struct {
	ID             int
	AccountID      int
	Account        Account
	OccuredAt      time.Time
	StandartQty    int
	GlossQty       int
	PosterQty      int
	Total          int
	StandartAmtUsd float64
	GlossAmtUsd    float64
	PosterAmtUsd   float64
	TotalAmtUsd    float64
}
