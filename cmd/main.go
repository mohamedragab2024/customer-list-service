package main

import (
	"customer-list-service/internal/models"
	"customer-list-service/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	models.DB, err = gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	models.DB.AutoMigrate(&models.Customer{})

	r := gin.Default()

	routes.RegisterCustomerRoutes(r)

	r.Run(":8080")
}
