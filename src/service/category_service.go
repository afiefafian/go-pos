package service

import (
	"time"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type CategoryService struct {
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryService(categoryService *repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryService}
}

func (s *CategoryService) FindAll(params *model.PaginationQuery) ([]model.GetCategoryResponse, *model.PaginationResponse, error) {
	// Get data
	categories, err := s.CategoryRepository.FindAll(params)
	if err != nil {
		return nil, nil, err
	}

	// Get total
	var count int
	count, err = s.CategoryRepository.Count()
	if err != nil {
		return nil, nil, err
	}

	// Format data
	var categoriesResponse = make([]model.GetCategoryResponse, 0)
	for _, cashier := range categories {
		categoriesResponse = append(categoriesResponse, model.GetCategoryResponse{
			ID:   cashier.ID,
			Name: cashier.Name,
		})
	}

	pagination := &model.PaginationResponse{
		Total: count,
		Skip:  params.Skip,
		Limit: params.Limit,
	}

	return categoriesResponse, pagination, nil
}

func (s *CategoryService) GetByID(id int64) (*model.GetCategoryResponse, error) {
	cashier, err := s.CategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &model.GetCategoryResponse{
		ID:   cashier.ID,
		Name: cashier.Name,
	}, nil
}

func (s *CategoryService) Create(request model.CreateCategoryRequest) (*model.CreateCategoryResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	currentTime := time.Now()
	category := entity.Category{
		Name:      request.Name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	id, err := s.CategoryRepository.Create(category)
	if err != nil {
		return nil, err
	}

	return &model.CreateCategoryResponse{
		ID:        id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (s *CategoryService) UpdateByID(request model.UpdateCategoryRequest) error {
	if _, err := s.CategoryRepository.GetByID(request.ID); err != nil {
		return err
	}

	if err := helper.ValidateStruct(request); err != nil {
		return err
	}

	category := entity.Category{
		ID:   request.ID,
		Name: request.Name,
	}

	if err := s.CategoryRepository.UpdateByID(category); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteByID(id int64) error {
	if _, err := s.CategoryRepository.GetByID(id); err != nil {
		return err
	}

	if err := s.CategoryRepository.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
