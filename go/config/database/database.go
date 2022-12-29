package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	ID         int
	Name       string
	Website    string
	Lat        float64
	Long       float64
	PrimaryPoc string
	SalesRepID int
	SalesRep   SalesRep
}

type WebEvent struct {
	ID        int
	AccountID int
	Account   Account
	OccuredAt time.Time
	Channel   string
}

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

type SalesRep struct {
	ID       int
	Name     string
	RegionID int
	Region   Region
}

type Region struct {
	ID   int
	Name string
}

func New() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=web_events port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Account{}, &WebEvent{}, &Order{}, &SalesRep{}, &Region{})

	return db
}
