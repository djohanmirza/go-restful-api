package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.ProductResponse{}, err
	}

	product := domain.Product{
		//ProductID:   request.ProductID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		StockQty:    request.StockQty,
		Category:    request.Category,
		SKU:         request.SKU,
		TaxRate:     request.TaxRate,
	}

	savedProduct, err := service.ProductRepository.Save(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(savedProduct), nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.ProductResponse{}, err
	}

	product, err := service.ProductRepository.FindById(ctx, request.ProductID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.StockQty = request.StockQty
	product.Category = request.Category
	product.SKU = request.SKU
	product.TaxRate = request.TaxRate

	updatedProduct, err := service.ProductRepository.Update(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(updatedProduct), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) error {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return err
	}

	return service.ProductRepository.Delete(ctx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) (web.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(product), nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToProductResponses(products), nil
}
