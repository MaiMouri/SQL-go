package usecase

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"
)

type WebEventUseCase interface {
	GetWebEvent(id int) (result *model.WebEvent, err error)
	GetWebEvents() (result []model.WebEvent, err error)
	// CreateCustomer(name string, age int) error
	// UpdateCustomer(id int, name string, age int) error
	// DeleteCustomer(id int) error
}

type webEventUseCase struct {
	webEventRepository repository.WebEventRepository
}

func NewWebEventUseCase(wr repository.WebEventRepository) WebEventUseCase {
	return &webEventUseCase{
		webEventRepository: wr,
	}
}

func (wu *webEventUseCase) GetWebEvent(id int) (result *model.WebEvent, err error) {
	webEvent, err := wu.webEventRepository.GetWebEvent(id)
	if err != nil {
		return nil, err
	}

	return webEvent, nil
}

func (wu *webEventUseCase) GetWebEvents() (result []model.WebEvent, err error) {
	webEvents, err := wu.webEventRepository.GetWebEvents()
	if err != nil {
		return nil, err
	}

	return webEvents, nil
}
