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

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.ProductCreateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.ProductCreateRequest{
				Name:        "Laptop",
				Description: "High-end laptop",
				Price:       1000,
				StockQty:    10,
				CategoryID:  1,
				SKU:         "LPT123",
				TaxRate:     0.1,
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{
					ProductID:   "1",
					Name:        "Laptop",
					Description: "High-end laptop",
					Price:       1000,
					StockQty:    10,
					CategoryId:  1,
					SKU:         "LPT123",
					TaxRate:     0.1,
				}, nil)
			},
			expect: web.ProductResponse{
				ProductID:   "1",
				Name:        "Laptop",
				Description: "High-end laptop",
				Price:       1000,
				StockQty:    10,
				CategoryID:  1,
				SKU:         "LPT123",
				TaxRate:     0.1,
			},
			expectErr: false,
		},
		{
			name: "validation error",
			input: web.ProductCreateRequest{
				Name: "",
			},
			mock:      func() {},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
		{
			name: "repository error",
			input: web.ProductCreateRequest{
				Name:        "Smartphone",
				Description: "Flagship phone",
				Price:       700,
				StockQty:    20,
				CategoryID:  2,
				SKU:         "SPH456",
				TaxRate:     0.1,
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("database error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
