package usecase

import (
	"sql-go/go/domain/model"
	"sql-go/go/domain/repository"
)

type RegionUseCase interface {
	GetRegion(id int) (result *model.Region, err error)
	GetRegions() (result []model.Region, err error)
	// CreateCustomer(name string, age int) error
	// UpdateCustomer(id int, name string, age int) error
	// DeleteCustomer(id int) error
}

type regionUseCase struct {
	regionRepository repository.RegionRepository
}

func NewRegionUseCase(rr repository.RegionRepository) RegionUseCase {
	return &regionUseCase{
		regionRepository: rr,
	}
}

func (au *regionUseCase) GetRegion(id int) (result *model.Region, err error) {
	region, err := au.regionRepository.GetRegion(id)
	if err != nil {
		return nil, err
	}

	return region, nil
}

func (au *regionUseCase) GetRegions() (result []model.Region, err error) {
	regions, err := au.regionRepository.GetRegions()
	if err != nil {
		return nil, err
	}

	return regions, nil
}
