package handlers

import (
	"customer-list-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	term := c.Query("term")

	if term != "" {
		// Search by name or email if term exists
		models.DB.Where("name LIKE @term OR email LIKE @term", map[string]interface{}{"term": "%" + term + "%"}).Find(&customers)
	} else {
		// Get all customers if no term provided
		models.DB.Find(&customers)
	}

	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c *gin.Context) {
	var customer models.Customer
	if err := models.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := models.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customer
	if err := models.DB.First(&customer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	models.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
