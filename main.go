package main

import (
	"github.com/afiffahreza/dorayaki-backend/connection"
	"github.com/afiffahreza/dorayaki-backend/controller"
	"github.com/afiffahreza/dorayaki-backend/repository"
	"github.com/afiffahreza/dorayaki-backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = connection.SetupDatabaseConnection()
	tokoRepository     repository.TokoRepository     = repository.NewTokoRepository(db)
	tokoService        service.TokoService           = service.NewTokoService(tokoRepository)
	tokoController     controller.TokoController     = controller.NewTokoController(tokoService)
	dorayakiRepository repository.DorayakiRepository = repository.NewDorayakiRepository(db)
	dorayakiService    service.DorayakiService       = service.NewDorayakiService(dorayakiRepository)
	dorayakiController controller.DorayakiController = controller.NewDorayakiController(dorayakiService)
	stockRepository    repository.StockRepository    = repository.NewStockRepository(db)
	stockService       service.StockService          = service.NewStockService(stockRepository)
	stockController    controller.StockController    = controller.NewStockController(stockService)
)

func main() {
	defer connection.CloseDatabaseConnection(db)
	r := gin.Default()

	/*
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Hello, World!"})
		})
	*/

	tokoRoutes := r.Group("api/toko")
	{
		tokoRoutes.GET("/", tokoController.All)
		tokoRoutes.POST("/", tokoController.Insert)
		tokoRoutes.PUT("/:id", tokoController.Update)
		tokoRoutes.DELETE("/:id", tokoController.Delete)
		tokoRoutes.GET("/:id", tokoController.FindByID)
		tokoRoutes.GET("/query", tokoController.FindTokoQuery)
	}

	dorayakiRoutes := r.Group("api/dorayaki")
	{
		dorayakiRoutes.GET("/", dorayakiController.All)
		dorayakiRoutes.GET("/:id", dorayakiController.FindByID)
		dorayakiRoutes.POST("/", dorayakiController.Insert)
		dorayakiRoutes.PUT("/:id", dorayakiController.Update)
		dorayakiRoutes.DELETE("/:id", dorayakiController.Delete)
	}

	stockRoutes := r.Group("api/stock")
	{
		stockRoutes.GET("/", stockController.All)
		stockRoutes.POST("/", stockController.Insert)
		stockRoutes.PUT("/:id", stockController.Update)
		stockRoutes.DELETE("/:id", stockController.Delete)
	}

	r.Run()
}
