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
	db             *gorm.DB                  = connection.SetupDatabaseConnection()
	tokoRepository repository.TokoRepository = repository.NewTokoRepository(db)
	tokoService    service.TokoService       = service.NewTokoService(tokoRepository)
	tokoController controller.TokoController = controller.NewTokoController(tokoService)
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
		tokoRoutes.GET("/kecamatan/:kec", tokoController.FindTokoKecamatan)
		tokoRoutes.GET("/provinsi/:prov", tokoController.FindTokoProvinsi)
	}

	r.Run()
}
