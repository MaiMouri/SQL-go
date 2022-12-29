package usecase

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"
)

type OrderUseCase interface {
	GetOrder(id int) (result *model.Order, err error)
	GetOrders() (result []model.Order, err error)
	// CreateCustomer(name string, age int) error
	// UpdateCustomer(id int, name string, age int) error
	// DeleteCustomer(id int) error
}

type orderUseCase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUseCase(or repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		orderRepository: or,
	}
}

func (ou *orderUseCase) GetOrder(id int) (result *model.Order, err error) {
	order, err := ou.orderRepository.GetOrder(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (ou *orderUseCase) GetOrders() (result []model.Order, err error) {
	orders, err := ou.orderRepository.GetOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
