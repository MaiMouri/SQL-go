package controller

import (
	"fmt"
	"sql-go/go/usecase"

	"github.com/gin-gonic/gin"
)

type accountController struct {
	accountUseCase usecase.AccountUseCase
}

func NewAccountController(au usecase.AccountUseCase) accountController {
	return accountController{
		accountUseCase: au,
	}

}

func (ac *accountController) Index(c *gin.Context) {

	accounts, err := ac.accountUseCase.GetAccounts()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"accounts": accounts})
}

// func GetAll(c *gin.Context) ([]database.Account, error){
// 	var account database.Account
// 	accounts := db.Find(&account)
// 	fmt.Println(accounts)
// 	c.JSON(200, gin.H{"accounts": accounts})
// }
