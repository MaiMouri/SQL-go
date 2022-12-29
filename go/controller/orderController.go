package controller

import (
	"fmt"
	"sql-go/go/usecase"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(au usecase.OrderUseCase) orderController {
	return orderController{
		orderUseCase: au,
	}

}

func (ac *orderController) Index(c *gin.Context) {

	orders, err := ac.orderUseCase.GetOrders()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"orders": orders})
}

// func GetAll(c *gin.Context) ([]database.Order, error){
// 	var order database.Order
// 	orders := db.Find(&order)
// 	fmt.Println(orders)
// 	c.JSON(200, gin.H{"orders": orders})
// }
