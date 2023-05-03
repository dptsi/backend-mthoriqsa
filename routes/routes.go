// routes/routes.go
package routes

import (
	"cobagolang/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/customers", controllers.FindCustomers)
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers/:id", controllers.FindCustomer)
	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("customers/:id", controllers.DeleteCustomer)
	return r
}
