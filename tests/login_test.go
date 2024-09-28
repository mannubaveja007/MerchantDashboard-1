package controllers_test

import (
	"bytes"
	"encoding/json"
	"merchant-dashboard/controllers"
	"merchant-dashboard/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	product := models.Product{
		MerchantID: 123,
		ProductID:  456,
		Name:       "Test Product",
		Price:      99,
		Quantity:   10,
	}

	body, _ := json.Marshal(product)
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controllers.CreateProduct(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Product created")
}

func TestGetProducts_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest(http.MethodGet, "/products?merchant_id=123", nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controllers.GetProducts(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var products []models.Product
	err := json.Unmarshal(w.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}

func TestUpdateProduct_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	invalidJSON := `{"MerchantID": "123", "ProductID": "456", "Name": "Test Product", "Price": "invalid", "Quantity": 10}`
	req, _ := http.NewRequest(http.MethodPut, "/products", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controllers.UpdateProduct(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
