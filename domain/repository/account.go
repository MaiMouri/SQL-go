package repository

import (
	"sql-go/domain/model"
)

type AccountRepository interface {
	GetAccount(id int) (result *model.Account, err error)
	GetAccounts() (result []model.Account, err error)
	// Create(c model.Customer) error
	// Update(c model.Customer) error
	// Delete(c model.Customer) error
}
