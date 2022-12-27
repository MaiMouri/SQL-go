package persistance

import (
	"sql-go/domain/model"
	"sql-go/domain/repository"

	"gorm.io/gorm"
)

type accountPersistance struct {
	Conn *gorm.DB
}

func NewAccountPersistance(conn *gorm.DB, a repository.AccountRepository) *accountPersistance {
	return &accountPersistance{Conn: conn}
}

func (cr *accountPersistance) GetAccount(id int) (result *model.Account, err error) {

	var account model.Account
	if result := cr.Conn.First(&account, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &account, nil
}

func (cr *accountPersistance) GetAccounts() (result []model.Account, err error) {

	var accounts []model.Account
	if result := cr.Conn.Find(&accounts); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return accounts, nil
}
