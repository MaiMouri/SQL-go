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
	SalesRepId int
}

type WebEvents struct {
	ID        int
	AccountID int
	Account   Account
	OccuredAt time.Time
	Age       int
	Channel   string
}

type Orders struct {
	ID             int
	AccountID      int
	Account        Account
	StandartQty    string
	PosterQty      string
	Total          int
	OccuredAt      time.Time
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

	db.AutoMigrate(&Account{}, &WebEvents{}, &Orders{}, &SalesRep{}, &Region{})

	return db
}
