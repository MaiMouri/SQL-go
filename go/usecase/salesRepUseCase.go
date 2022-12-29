package usecase

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"
)

type SalesRepUseCase interface {
	GetSalesRep(id int) (result *model.SalesRep, err error)
	GetSalesReps() (result []model.SalesRep, err error)
	// CreateCustomer(name string, age int) error
	// UpdateCustomer(id int, name string, age int) error
	// DeleteCustomer(id int) error
}

type salesRepUseCase struct {
	salesRepRepository repository.SalesRepRepository
}

func NewSalesRepUseCase(sr repository.SalesRepRepository) SalesRepUseCase {
	return &salesRepUseCase{
		salesRepRepository: sr,
	}
}

func (au *salesRepUseCase) GetSalesRep(id int) (result *model.SalesRep, err error) {
	salesRep, err := au.salesRepRepository.GetSalesRep(id)
	if err != nil {
		return nil, err
	}

	return salesRep, nil
}

func (au *salesRepUseCase) GetSalesReps() (result []model.SalesRep, err error) {
	salesReps, err := au.salesRepRepository.GetSalesReps()
	if err != nil {
		return nil, err
	}

	return salesReps, nil
}
