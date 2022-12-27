package usecase

import (
	"sql-go/domain/model"
	"sql-go/domain/repository"
)

type AccountUseCase interface {
	GetAccount(id int) (result *model.Account, err error)
	GetAccounts() (result []model.Account, err error)
	// CreateCustomer(name string, age int) error
	// UpdateCustomer(id int, name string, age int) error
	// DeleteCustomer(id int) error
}

type accountUseCase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUseCase(ar repository.AccountRepository) AccountUseCase {
	return &accountUseCase{
		accountRepository: ar,
	}
}

func (au *accountUseCase) GetAccount(id int) (result *model.Account, err error) {
	account, err := au.accountRepository.GetAccount(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (au *accountUseCase) GetAccounts() (result []model.Account, err error) {
	accounts, err := au.accountRepository.GetAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
