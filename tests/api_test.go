package tests

import (
	"bytes"
	"ecommerce-api/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterUser(t *testing.T) {
	r := gin.Default()
	r.POST("/users/register", controllers.Register) // 使用 controllers.Register

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
