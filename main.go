package main

import (
	"fmt"
	// "cobagolang/models"
	// "cobagolang/routes"
)

func main() {
	arr := [4]int{1, 2, 3}
	arr[2] = 4
	fmt.Print(arr)
	// db := models.SetupDB()
	// db.AutoMigrate(&models.Customer{})

	// r := routes.SetupRoutes(db)
	// r.Run()
}
