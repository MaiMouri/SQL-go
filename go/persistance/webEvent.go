package persistance

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"

	"gorm.io/gorm"
)

type webEventPersistance struct {
	Conn *gorm.DB
}

func NewWebEventPersistance(conn *gorm.DB, a repository.WebEventRepository) *webEventPersistance {
	return &webEventPersistance{Conn: conn}
}

func (cr *webEventPersistance) GetWebEvent(id int) (result *model.WebEvent, err error) {

	var webEvent model.WebEvent
	if result := cr.Conn.First(&webEvent, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &webEvent, nil
}

func (cr *webEventPersistance) GetWebEvents() (result []model.WebEvent, err error) {

	var webEvents []model.WebEvent
	if result := cr.Conn.Preload("Account").Find(&webEvents); result.Error != nil {
		// if result := cr.Conn.Preload("Account").Preload("SalesRep").Find(&webEvents); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return webEvents, nil
}
