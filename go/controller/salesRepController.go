package controller

import (
	"fmt"
	"sql-go/go/usecase"

	"github.com/gin-gonic/gin"
)

type salesRepController struct {
	salesRepUseCase usecase.SalesRepUseCase
}

func NewSalesRepController(au usecase.SalesRepUseCase) salesRepController {
	return salesRepController{
		salesRepUseCase: au,
	}

}

func (ac *salesRepController) Index(c *gin.Context) {

	salesReps, err := ac.salesRepUseCase.GetSalesReps()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"salesReps": salesReps})
}

// func GetAll(c *gin.Context) ([]database.SalesRep, error){
// 	var salesRep database.SalesRep
// 	salesReps := db.Find(&salesRep)
// 	fmt.Println(salesReps)
// 	c.JSON(200, gin.H{"salesReps": salesReps})
// }
