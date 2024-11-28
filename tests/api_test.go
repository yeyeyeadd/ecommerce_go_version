package tests

import (
	"bytes"
	"ecommerce-api/controllers"
	"ecommerce-api/routes"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterUser(t *testing.T) {
	r := gin.Default()
	r.POST("/users/register", controllers.Register) // controllers.Register

	body := []byte(`{"username":"testuser", "email":"test@example.com", "password":"password123"}`)
	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d", w.Code)
	}
}

func TestUserLogin(t *testing.T) {
	r := gin.Default()
	r.POST("/users/login", controllers.Login)

	body := []byte(`{"email":"test@example.com", "password":"password123"}`)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", w.Code)
	}

	responseData := w.Body.String()
	if !strings.Contains(responseData, "token") {
		t.Fatalf("Expected token in response but got %s", responseData)
	}
}

func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	var response []map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	assert.Greater(t, len(response), 0, "Expected non-empty product list")
}

func TestCreateProduct(t *testing.T) {
	payload := map[string]interface{}{
		"name":        "New Product",
		"description": "A test product",
		"price":       99.99,
		"stock":       100,
		"seller_id":   1,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test_valid_token") // Replace with valid token

	rr := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201")
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	assert.Equal(t, "New Product", response["name"])
}
