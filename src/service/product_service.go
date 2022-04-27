package service

import (
	"time"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(paymentRepository *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: paymentRepository}
}

func (s *ProductService) FindAll(params model.GetProductQuery) ([]model.GetProductResponse, *model.PaginationResponse, error) {
	// Get data
	products, err := s.ProductRepository.FindAll(params)
	if err != nil {
		return nil, nil, err
	}

	// Get total
	var count int
	count, err = s.ProductRepository.Count(params)
	if err != nil {
		return nil, nil, err
	}

	// Format data
	var productsResponse = make([]model.GetProductResponse, 0)
	for _, product := range products {
		productResponse := model.GetProductResponse{
			ID:    product.ID,
			SKU:   product.SKU(),
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
			Image: product.Image,
		}

		if product.Category != nil {
			productResponse.Category = &model.GetCategoryResponse{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			}
		}

		if product.Discount != nil {
			productResponse.Discount = &model.ProductDiscountResponse{
				ID:              product.Discount.ID,
				Qty:             product.Discount.Qty,
				Type:            product.Discount.Type,
				Result:          product.Discount.Result,
				ExpiredAt:       product.Discount.ExpiredAtDate(),
				ExpiredAtFormat: product.Discount.ExpiredAtFormat(),
				StringFormat:    product.Discount.StringFormat(),
			}
		}

		productsResponse = append(productsResponse, productResponse)
	}

	pagination := &model.PaginationResponse{
		Total: count,
		Skip:  params.Pagination.Skip,
		Limit: params.Pagination.Limit,
	}

	return productsResponse, pagination, nil
}

func (s *ProductService) GetByID(id int64) (*model.GetProductResponse, error) {
	product, err := s.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	productResponse := &model.GetProductResponse{
		ID:    product.ID,
		SKU:   product.SKU(),
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
		Image: product.Image,
	}

	if product.Category != nil {
		productResponse.Category = &model.GetCategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		}
	}

	if product.Discount != nil {
		productResponse.Discount = &model.ProductDiscountResponse{
			ID:              product.Discount.ID,
			Qty:             product.Discount.Qty,
			Type:            product.Discount.Type,
			Result:          product.Discount.Result,
			ExpiredAt:       product.Discount.ExpiredAtDate(),
			ExpiredAtFormat: product.Discount.ExpiredAtFormat(),
			StringFormat:    product.Discount.StringFormat(),
		}
	}

	return productResponse, nil
}

func (s *ProductService) Create(request model.CreateProductRequest) (*model.CreateProductResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	currentTime := time.Now()
	product := entity.Product{
		Name:  request.Name,
		Image: request.Image,
		Stock: request.Stock,
		Price: request.Price,
		Category: &entity.Category{
			ID: request.CategoryID,
		},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if request.Discount != nil {
		product.Discount = &entity.ProductDiscount{
			Type:      request.Discount.Type,
			Qty:       request.Discount.Qty,
			Result:    request.Discount.Result,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		}
	}

	id, err := s.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	product.ID = id

	return &model.CreateProductResponse{
		ID:         id,
		CategoryID: request.CategoryID,
		SKU:        product.SKU(),
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
		Image:      product.Image,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}, nil
}

func (s *ProductService) UpdateByID(request model.UpdateProductRequest) error {
	if _, err := s.ProductRepository.GetByID(request.ID); err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	product := entity.Product{
		ID:    request.ID,
		Name:  request.Name,
		Image: request.Image,
		Stock: request.Stock,
		Price: request.Price,
		Category: &entity.Category{
			ID: request.CategoryID,
		},
	}

	if err := s.ProductRepository.UpdateByID(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteByID(id int64) error {
	if _, err := s.ProductRepository.GetByID(id); err != nil {
		return err
	}

	if err := s.ProductRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
