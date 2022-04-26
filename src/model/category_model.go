package model

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type UpdateCategoryRequest struct {
	ID   int64
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type CreateCategoryResponse struct {
	ID        int64     `json:"categoryId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetCategoryResponse struct {
	ID   int64  `json:"categoryId"`
	Name string `json:"name"`
}
