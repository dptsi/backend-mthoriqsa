package main

import (
	//"cobagolang/functions"
	//"fmt"
	"cobagolang/config"
	"cobagolang/models"

	"cobagolang/routes"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&models.Customer{})

	r := routes.SetupRoutes(db)
	r.Run(":8032")
}
