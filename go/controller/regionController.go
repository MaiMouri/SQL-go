package controller

import (
	"fmt"
	"sql-go/go/usecase"

	"github.com/gin-gonic/gin"
)

type regionController struct {
	regionUseCase usecase.RegionUseCase
}

func NewRegionController(au usecase.RegionUseCase) regionController {
	return regionController{
		regionUseCase: au,
	}

}

func (ac *regionController) Index(c *gin.Context) {

	regions, err := ac.regionUseCase.GetRegions()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"regions": regions})
}

// func GetAll(c *gin.Context) ([]database.Region, error){
// 	var region database.Region
// 	regions := db.Find(&region)
// 	fmt.Println(regions)
// 	c.JSON(200, gin.H{"regions": regions})
// }
