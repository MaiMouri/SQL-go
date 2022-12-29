package persistance

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"

	"gorm.io/gorm"
)

type orderPersistance struct {
	Conn *gorm.DB
}

func NewOrderPersistance(conn *gorm.DB, a repository.OrderRepository) *orderPersistance {
	return &orderPersistance{Conn: conn}
}

func (cr *orderPersistance) GetOrder(id int) (result *model.Order, err error) {

	var order model.Order
	if result := cr.Conn.First(&order, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &order, nil
}

func (cr *orderPersistance) GetOrders() (result []model.Order, err error) {

	var orders []model.Order
	if result := cr.Conn.Find(&orders); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return orders, nil
}
