package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	mockValidator := validator.New()
	employeeService := NewEmployeeService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.EmployeeCreateRequest
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.EmployeeCreateRequest{
				Name: "John Doe", Role: "Developer", Email: "john@example.com", Phone: "1234567890", DateHired: "2025-01-01",
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{
					EmployeeID: "1", Name: "John Doe", Role: "Developer", Email: "john@example.com", Phone: "1234567890", DateHired: "2025-01-01",
				}, nil)
			},
			expect: web.EmployeeResponse{
				EmployeeID: "1", Name: "John Doe", Role: "Developer", Email: "john@example.com", Phone: "1234567890", DateHired: "2025-01-01",
			},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.EmployeeCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestFindByIdEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := NewEmployeeService(mockRepo, validator.New())

	tests := []struct {
		name      string
		id        string
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name: "success",
			id:   "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Employee{
					EmployeeID: "1", Name: "John Doe", Role: "Developer",
				}, nil)
			},
			expect: web.EmployeeResponse{
				EmployeeID: "1", Name: "John Doe", Role: "Developer",
			},
			expectErr: false,
		},
		{
			name: "not found",
			id:   "99",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "99").Return(domain.Employee{}, errors.New("not found"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.FindById(context.Background(), tt.id)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
