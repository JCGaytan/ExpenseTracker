package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"expense-tracker/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	config.LoadConfig()
	config.InitDB()
	config.SeedUsers()

	r := gin.Default()
	r.POST("/auth/login", Login)

	loginPayload := map[string]string{
		"email":    "test1@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(loginPayload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	_, ok := resp["token"]
	assert.True(t, ok, "token should be present in response")
}
