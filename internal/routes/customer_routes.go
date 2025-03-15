package routes

import (
	"customer-list-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	r.POST("/customers", handlers.CreateCustomer)
	r.GET("/customers/:id", handlers.GetCustomer)
	r.GET("/customers", handlers.GetCustomers)
	r.PUT("/customers/:id", handlers.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)
}
