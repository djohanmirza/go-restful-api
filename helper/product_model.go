package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
		Category:    product.Category,
		SKU:         product.SKU,
		TaxRate:     product.TaxRate,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
