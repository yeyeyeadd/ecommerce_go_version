package tests

import (
	"bytes"
	"ecommerce-api/controllers"
	"ecommerce-api/models"
	"ecommerce-api/routes"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	err := os.Chdir("/Users/yansongwang/Desktop/untitled folder/ecommerce-api") // replace with your root directory
	if err != nil {
		log.Fatalf("Error changing working directory: %v", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file22: %v", err)
	}

	// Initialize db
	models.InitDB()

	// Test run
	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	router := routes.SetupRouter()

	r := gin.Default()
	r.POST("/users/register", controllers.Register) // 使用 controllers.Register

	body := []byte(`{"username":"test4", "email":"test4@qq.com", "password":"123456"}`)
	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	// 打印调试信息
	t.Logf("Response status: %d", rr.Code)
	t.Logf("Raw response: %s", rr.Body.String())

	// 校验响应状态码
	if rr.Code != http.StatusCreated {
		t.Fatalf("Expected status code 200, got %d", rr.Code)
	}

	// 解析响应数据
	var responseData map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &responseData)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	t.Logf("Response data: %v", responseData)
}

func TestUserLogin(t *testing.T) {
	router := routes.SetupRouter()

	// Check JSON format
	reqBody := `{"email":"test4@qq.com","password":"123456"}`
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Test recorder
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// print info
	t.Logf("Request body: %s", reqBody)
	t.Logf("Response status: %d", rr.Code)
	t.Logf("Raw response: %s", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", rr.Code)
	}

	// Parse the response data.
	var responseData map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &responseData)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	t.Logf("Response data: %v", responseData)

	// Validate the response content
	if responseData["message"] != "Login successful" {
		t.Errorf("Expected message 'Login successful', got %v", responseData["message"])
	}
	if _, exists := responseData["token"]; !exists {
		t.Errorf("Token not found in response")
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
	t.Logf("Response data: %v", response)

	assert.Greater(t, len(response), 0, "Expected non-empty product list")
}

func TestCreateProduct(t *testing.T) {
	payload := map[string]interface{}{
		"name":        "Product2",
		"description": "A test product",
		"price":       9.99,
		"stock":       100,
		"seller_id":   2,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Replace with valid token
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI3OTIxMjUsInVzZXJfaWQiOjJ9.yM3X6CdoxTTDTrEaod8aqnu5-EpSct2fQniXAwfIo54")

	rr := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201")
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	t.Logf("Response data: %v", response)

	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	assert.Equal(t, "Product2", response["name"])
}

func TestCreateOrder(t *testing.T) {
	// 假设购买者的 user_id 为 1
	// buyerID := 1

	payload := map[string]interface{}{
		"items": []map[string]interface{}{
			{"product_id": 61, "quantity": 2},
			{"product_id": 62, "quantity": 1},
		},
	}
	body, _ := json.Marshal(payload)

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI3OTIxMjUsInVzZXJfaWQiOjJ9.yM3X6CdoxTTDTrEaod8aqnu5-EpSct2fQniXAwfIo54")

	// 创建响应记录器
	rr := httptest.NewRecorder()

	// 使用 SetupRouter 初始化路由
	router := routes.SetupRouter()

	// 模拟请求
	router.ServeHTTP(rr, req)

	// 验证响应
	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201")
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	t.Logf("Response data: %v", response)

	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 确保返回的订单 ID 大于 0
	assert.Greater(t, response["id"].(float64), 0, "Order ID should be greater than 0")
}
