package persistance

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"

	"gorm.io/gorm"
)

type regionPersistance struct {
	Conn *gorm.DB
}

func NewRegionPersistance(conn *gorm.DB, a repository.RegionRepository) *regionPersistance {
	return &regionPersistance{Conn: conn}
}

func (cr *regionPersistance) GetRegion(id int) (result *model.Region, err error) {

	var region model.Region
	if result := cr.Conn.First(&region, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &region, nil
}

func (cr *regionPersistance) GetRegions() (result []model.Region, err error) {

	var regions []model.Region
	if result := cr.Conn.Find(&regions); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return regions, nil
}
