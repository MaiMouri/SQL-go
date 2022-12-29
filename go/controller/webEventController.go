package controller

import (
	"fmt"
	"sql-go/go/usecase"

	"github.com/gin-gonic/gin"
)

type webEventController struct {
	webEventUseCase usecase.WebEventUseCase
}

func NewWebEventController(au usecase.WebEventUseCase) webEventController {
	return webEventController{
		webEventUseCase: au,
	}

}

func (ac *webEventController) Index(c *gin.Context) {

	webEvents, err := ac.webEventUseCase.GetWebEvents()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"webEvents": webEvents})
}

// func GetAll(c *gin.Context) ([]database.WebEvent, error){
// 	var webEvent database.WebEvent
// 	webEvents := db.Find(&webEvent)
// 	fmt.Println(webEvents)
// 	c.JSON(200, gin.H{"webEvents": webEvents})
// }
