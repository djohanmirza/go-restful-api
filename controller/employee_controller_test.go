package controller

import (
	"bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupEmployeeTestApp(mockService *mocks.MockEmployeeService) *fiber.App {
	app := fiber.New()
	employeeController := NewEmployeeController(mockService)

	api := app.Group("/api")
	employees := api.Group("/employees")
	employees.Post("/", employeeController.Create)
	employees.Put("/:employeeId", employeeController.Update)
	employees.Delete("/:employeeId", employeeController.Delete)
	employees.Get("/:employeeId", employeeController.FindById)
	employees.Get("/", employeeController.FindAll)

	return app
}

func TestEmployeeController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockEmployeeService(ctrl)
	app := setupEmployeeTestApp(mockService)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		setupMock      func()
		expectedStatus int
		expectedBody   web.WebResponse
	}{
		{
			name:   "Create employee - success",
			method: "POST",
			url:    "/api/employees/",
			body:   web.EmployeeCreateRequest{Name: "John Doe", Role: "Engineer"},
			setupMock: func() {
				mockService.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(web.EmployeeResponse{EmployeeID: "1", Name: "John Doe", Role: "Engineer"}, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: web.WebResponse{
				Code:   http.StatusCreated,
				Status: "Created",
				Data:   web.EmployeeResponse{EmployeeID: "1", Name: "John Doe", Role: "Engineer"},
			},
		},
		{
			name:   "Find employee by ID - success",
			method: "GET",
			url:    "/api/employees/1",
			body:   nil,
			setupMock: func() {
				mockService.EXPECT().
					FindById(gomock.Any(), "1").
					Return(web.EmployeeResponse{EmployeeID: "1", Name: "John Doe", Role: "Engineer"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.EmployeeResponse{EmployeeID: "1", Name: "John Doe", Role: "Engineer"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			var reqBody []byte
			if tt.body != nil {
				reqBody, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, tt.url, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var respBody web.WebResponse
			json.NewDecoder(resp.Body).Decode(&respBody)

			if dataMap, ok := respBody.Data.(map[string]interface{}); ok {
				respBody.Data = web.EmployeeResponse{
					EmployeeID: dataMap["employee_id"].(string),
					Name:       dataMap["name"].(string),
					Role:       dataMap["role"].(string),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}
