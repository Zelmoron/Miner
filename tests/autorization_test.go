package test

import (
	"WebSocket/internal/endpoints"
	"WebSocket/internal/requests"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type MockServices struct{}

func (m *MockServices) Registration(user requests.UserRegRequest) error {
	type testStruct struct {
		email string
	}
	data := testStruct{
		email: "john.doe@example.com",
	}
	if user.Email != data.email {
		return errors.New("")
	}
	return nil
}

func (m *MockServices) Login(user requests.UserLoginRequest) (string, string, error) {
	type testStruct struct {
		email string
	}
	data := testStruct{
		email: "john.doe@example.com",
	}
	if user.Email != data.email {
		return "", "", errors.New("")
	}
	return "token", "token", nil
}

func (m *MockServices) NewJWT(interface{}) (string, error) {
	return "", nil
}

func (m MockServices) Delete(string) error {
	return nil
}

func TestRegistration(t *testing.T) {
	app := fiber.New()
	mockServices := &MockServices{}
	endpoints := endpoints.New(mockServices)

	app.Post("/registration", endpoints.Registration)

	tests := []struct {
		name             string
		body             requests.UserRegRequest
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		{
			name: "Successful Registration",
			body: requests.UserRegRequest{
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: "securepassword123",
			},
			expectedStatus: http.StatusOK,
			expectedResponse: map[string]interface{}{
				"status": "OK",
			},
		},
		{
			name: "Bad Request - Validation Error",
			body: requests.UserRegRequest{
				Name:     "",
				Email:    "invalid-email",
				Password: "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"status": "BadRequest - Validation error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/registration", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.Equal(t, tt.expectedResponse, responseBody)
		})
	}
}

func TestLogin(t *testing.T) {
	app := fiber.New()
	mockServices := &MockServices{}
	endpoints := endpoints.New(mockServices)

	app.Post("/login", endpoints.Login)

	tests := []struct {
		name             string
		body             requests.UserLoginRequest
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		{
			name: "Successful Login",
			body: requests.UserLoginRequest{
				Email:    "john.doe@example.com",
				Password: "securepassword123",
			},
			expectedStatus: http.StatusOK,
			expectedResponse: map[string]interface{}{
				"status": "OK", "access_token": "token",
			},
		},
		{
			name: "Bad Request - Validation Error",
			body: requests.UserLoginRequest{
				Email:    "invalid-email",
				Password: "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"status": "BadRequest - Validation error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.Equal(t, tt.expectedResponse, responseBody)
		})
	}
}
