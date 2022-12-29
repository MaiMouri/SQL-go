package persistance

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"

	"gorm.io/gorm"
)

type salesRepPersistance struct {
	Conn *gorm.DB
}

func NewSalesRepPersistance(conn *gorm.DB, a repository.SalesRepRepository) *salesRepPersistance {
	return &salesRepPersistance{Conn: conn}
}

func (cr *salesRepPersistance) GetSalesRep(id int) (result *model.SalesRep, err error) {

	var salesRep model.SalesRep
	if result := cr.Conn.First(&salesRep, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &salesRep, nil
}

func (cr *salesRepPersistance) GetSalesReps() (result []model.SalesRep, err error) {

	var salesReps []model.SalesRep
	if result := cr.Conn.Preload("Region").Find(&salesReps); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return salesReps, nil
}
