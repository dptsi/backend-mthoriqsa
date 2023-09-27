// routes/routes.go
package routes

import (
	"cobagolang/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		cors.New(config)
	})

	//routes
	r.GET("/customers", controllers.FindCustomers)
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers/:id", controllers.FindCustomer)
	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	r.GET("/stocks", controllers.GetStocks)
	r.POST("/stocks", controllers.CreateStock)
	r.PATCH("/stocks/:id", controllers.UpdateStock)
	r.DELETE("/stocks/:id", controllers.DeleteStock)
	r.POST("/stocks/get-by-name-or-price", controllers.FindStockByNameOrPrice)
	r.GET("/stocks/get-by-price-range/:min/:max", controllers.FindStockByPriceRange)

	return r
}
