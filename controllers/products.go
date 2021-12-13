package controllers

import (
	"musayann/gorm-playground/config"
	"musayann/gorm-playground/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductData struct {
	Name  string `form:"name" json:"name" xml:"name"  binding:"required"`
	Price uint   `form:"price" json:"price" xml:"price"  binding:"required"`
	Code  string `form:"code" json:"code" xml:"code"  binding:"required"`
}

func CreateProduct(c *gin.Context) {

	var json ProductData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.Db.Create(&models.Product{Code: json.Code, Price: uint64(json.Price), Name: json.Name})

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Create Product",
	})
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.Db.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Products",
		"data":    products,
	})
}
