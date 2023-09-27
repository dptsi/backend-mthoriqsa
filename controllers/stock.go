package controllers

import (
	"cobagolang/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PostRequest struct {
	Nama  string `json:"nama"`
	Harga string `json:"harga"`
}

func GetStocks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stocks []models.Stock
	param := c.Query("id")
	var res = db

	log.Println(param)

	//if param null, get all data
	if param == "" {
		res = db.Raw("SELECT * FROM stocks").Scan(&stocks)
		log.Println("mrono")
	} else {
		res = db.Raw("SELECT * FROM stocks WHERE id = ?", param).Scan(&stocks)
		log.Println("mrene")
	}

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	//set json response
	// c.JSON(http.StatusOK, gin.H{"data": stocks})
	c.JSON(http.StatusOK, stocks)
}

func GetStocksById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stocks []models.Stock
	res := db.Raw("SELECT * FROM stocks WHERE id = ?", c.Param("id")).Scan(&stocks)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stocks})
}

func CreateStock(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stock models.Stock

	//validate input
	if err := c.ShouldBindJSON(&stock); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//create task
	task := models.Stock{Nama: stock.Nama, Harga: stock.Harga, Qty: stock.Qty}
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateStock(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stock models.Stock

	//find data
	if err := db.Where("id = ? ", c.Param("id")).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	//validate input
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//set update
	var updatedInput models.Stock
	updatedInput.ID = stock.ID
	updatedInput.Nama = stock.Nama
	updatedInput.Harga = stock.Harga
	updatedInput.Qty = stock.Qty
	updatedInput.CreatedAt = stock.CreatedAt
	updatedInput.UpdatedAt = time.Now()

	if err := db.Model(&stock).Update(&updatedInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, updatedInput)
}

func DeleteStock(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stock models.Stock
	if err := db.Where("id = ?", c.Param("id")).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	db.Delete(&stock)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindStockByNameOrPrice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var request PostRequest
	var stocks []models.Stock

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := "SELECT * FROM stocks"
	if request.Nama != "" {
		query += " WHERE nama LIKE '%" + request.Nama + "%'"
	} else if request.Harga != "" {
		query += " WHERE harga >= " + request.Harga
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Null parameters"})
		return
	}

	res := db.Raw(query).Scan(&stocks)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func FindStockByPriceRange(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var stocks []models.Stock

	res := db.Raw("SELECT * FROM stocks WHERE harga >= ? AND harga <= ?", c.Param("min"), c.Param("max")).Scan(&stocks)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
