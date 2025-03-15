package handlers

import (
	"bytes"
	"customer-list-service/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	models.DB, _ = gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	models.DB.AutoMigrate(&models.Customer{})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/customers", CreateCustomer)
	r.GET("/customers/:id", GetCustomer)
	r.PUT("/customers/:id", UpdateCustomer)
	r.DELETE("/customers/:id", DeleteCustomer)
	return r
}

func TestCreateCustomer(t *testing.T) {
	r := setupRouter()

	customer := models.Customer{Name: "Test User", Address: "123 Test St", Status: "Active", Email: "test@example.com"}
	jsonValue, _ := json.Marshal(customer)
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCustomer(t *testing.T) {
	customer := models.Customer{Name: "Test User", Address: "123 Test St", Status: "Active", Email: "test@example.com"}
	models.DB.Create(&customer)

	r := setupRouter()
	req, _ := http.NewRequest("GET", "/customers/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCustomer(t *testing.T) {
	customer := models.Customer{Name: "Test User", Address: "123 Test St", Status: "Active", Email: "test@example.com"}
	models.DB.Create(&customer)

	r := setupRouter()
	customer.Name = "Updated User"
	jsonValue, _ := json.Marshal(customer)
	req, _ := http.NewRequest("PUT", "/customers/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCustomer(t *testing.T) {
	customer := models.Customer{Name: "Test User", Address: "123 Test St", Status: "Active", Email: "test@example.com"}
	models.DB.Create(&customer)

	r := setupRouter()
	req, _ := http.NewRequest("DELETE", "/customers/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
