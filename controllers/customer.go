// controllers/customer.go
package controllers

import (
	"cobagolang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCustomerInput struct {
	Name  string `json:"cust_name"`
	Phone string `json:"cust_phone"`
	DOB   string `json:"cust_dob"`
}

type UpdateCustomerInput struct {
	Name  string `json:"cust_name"`
	Phone string `json:"cust_phone"`
	DOB   string `json:"cust_dob"`
}

// GET /customer
// Get all customers
func FindCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// POST /customer
// Create new customer
func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	dob, _ := time.Parse(date, input.DOB)

	// Create task
	task := models.Customer{Name: input.Name, Phone: input.Phone, DOB: dob}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /customer/:id
// Find a customer
func FindCustomer(c *gin.Context) { // Get model if exist
	var customer models.Customer

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// PATCH /customer/:id
// Update a customer
func UpdateCustomer(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	dob, _ := time.Parse(date, input.DOB)

	var updatedInput models.Customer
	updatedInput.DOB = dob
	updatedInput.Name = input.Name
	updatedInput.Phone = input.Phone

	db.Model(&customer).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DELETE /customer/:id
// Delete a task
func DeleteCustomer(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
