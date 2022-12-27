package main

import (
	"sql-go/config/database"
	"sql-go/controller"
	"sql-go/domain/repository"
	"sql-go/persistance"
	"sql-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.New()
	connect, _ := db.DB()
	defer connect.Close()

	var accountRepository repository.AccountRepository
	accountPersistance := persistance.NewAccountPersistance(db, accountRepository)
	accountUseCase := usecase.NewAccountUseCase(accountPersistance)
	accountController := controller.NewAccountController(accountUseCase)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world",
		})
	})
	r.GET("/accounts", accountController.Index)
	r.Run()
}

// func accounts() (c *gin.Context) {
// 	accounts, err := db.Find(&account)
// }
