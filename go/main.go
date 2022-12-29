package main

import (
	"sql-go/go/config/database"
	"sql-go/go/controller"
	"sql-go/go/domain/repository"
	"sql-go/go/persistance"
	"sql-go/go/usecase"
	"time"

	"github.com/gin-contrib/cors"

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

	var orderRepository repository.OrderRepository
	orderPersistance := persistance.NewOrderPersistance(db, orderRepository)
	orderUseCase := usecase.NewOrderUseCase(orderPersistance)
	orderController := controller.NewOrderController(orderUseCase)

	var webEventRepository repository.WebEventRepository
	webEventPersistance := persistance.NewWebEventPersistance(db, webEventRepository)
	webEventUseCase := usecase.NewWebEventUseCase(webEventPersistance)
	webEventController := controller.NewWebEventController(webEventUseCase)

	var salesRepRepository repository.SalesRepRepository
	salesRepPersistance := persistance.NewSalesRepPersistance(db, salesRepRepository)
	salesRepUseCase := usecase.NewSalesRepUseCase(salesRepPersistance)
	salesRepController := controller.NewSalesRepController(salesRepUseCase)

	var regionRepository repository.RegionRepository
	regionPersistance := persistance.NewRegionPersistance(db, regionRepository)
	regionUseCase := usecase.NewRegionUseCase(regionPersistance)
	regionController := controller.NewRegionController(regionUseCase)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world",
		})
	})
	r.GET("/accounts", accountController.Index)
	r.GET("/orders", orderController.Index)
	r.GET("/web_events", webEventController.Index)
	r.GET("/sales_reps", salesRepController.Index)
	r.GET("/regions", regionController.Index)
	r.GET("/query", regionController.Index)
	r.Run()
}

// func accounts() (c *gin.Context) {
// 	accounts, err := db.Find(&account)
// }
