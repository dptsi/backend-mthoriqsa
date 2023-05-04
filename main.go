package main

import (
	"cobagolang/functions"
	"fmt"
	// "cobagolang/models"
	// "cobagolang/routes"
)

//func main() {
// db := models.SetupDB()
// db.AutoMigrate(&models.Customer{})

// r := routes.SetupRoutes(db)
// r.Run()
//}

func main() {
	arr := []int{1, 2, 3, 10, 19}
	total := functions.Sum(arr)

	fmt.Println(total)
}
