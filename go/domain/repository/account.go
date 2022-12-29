package repository

import (
	"sql-go/go/domain/model"
)

type AccountRepository interface {
	GetAccount(id int) (result *model.Account, err error)
	GetAccounts() (result []model.Account, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}

type OrderRepository interface {
	GetOrder(id int) (result *model.Order, err error)
	GetOrders() (result []model.Order, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}

type WebEventRepository interface {
	GetWebEvent(id int) (result *model.WebEvent, err error)
	GetWebEvents() (result []model.WebEvent, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}

type RegionRepository interface {
	GetRegion(id int) (result *model.Region, err error)
	GetRegions() (result []model.Region, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}

type SalesRepRepository interface {
	GetSalesRep(id int) (result *model.SalesRep, err error)
	GetSalesReps() (result []model.SalesRep, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}
