package repository

import (
	"database/sql"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) FindAll(params *model.PaginationQuery) ([]entity.Category, error) {
	rows, err := r.db.Query(
		"SELECT id, name, created_at, updated_at FROM categories LIMIT ? OFFSET ?",
		params.Limit,
		params.Skip,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *CategoryRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CategoryRepository) GetByID(id int64) (*entity.Category, error) {
	var category entity.Category
	row := r.db.QueryRow("SELECT id, name, created_at, updated_at FROM categories WHERE id = ?", id)
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.EntityNotFound("Category")
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category entity.Category) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO categories (name) VALUES (?)",
		category.Name,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *CategoryRepository) UpdateByID(category entity.Category) error {
	_, err := r.db.Exec(
		"UPDATE categories SET name = ? WHERE id = ?",
		category.Name,
		category.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteByID(id int64) error {
	if _, err := r.db.Exec(
		"DELETE FROM categories WHERE id = ?",
		id,
	); err != nil {
		return err
	}
	return nil
}
