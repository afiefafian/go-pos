package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepositoryy(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) FindAll(params model.GetProductQuery) ([]entity.Product, error) {
	var values []interface{}
	var where []string

	if params.CategoryId != nil {
		values = append(values, params.CategoryId)
		where = append(where, fmt.Sprintf("%s = ?", "p.category_id"))
	}

	if params.Q != nil {
		values = append(values, params.Q)
		where = append(where, fmt.Sprintf("%s LIKE %s", "p.name", "'?%'"))
	}

	values = append(values, params.Pagination.Limit)
	values = append(values, params.Pagination.Skip)

	query := `
		SELECT p.id, p.name, p.stock, p.price, p.image, p.created_at, p.updated_at, 
			p.category_id, c.name AS category_name, 
			p.discount_id, d.qty AS discount_qty, d.result AS discount_result, d.type AS discount_type, d.expired_at AS discount_exp_at 
		FROM products p 
		LEFT JOIN categories c ON p.category_id = c.id 
		LEFT JOIN product_discounts d ON p.discount_id = d.id
	`
	if len(where) > 0 {
		query += " WHERE "
		query += strings.Join(where, " AND ")
	}
	query += " LIMIT ? OFFSET ?"

	rows, err := r.db.Query(
		query,
		values...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product
		var category entity.Category
		var discount entity.ProductDiscount

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Stock,
			&product.Price,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,

			&category.ID,
			&category.Name,

			&discount.ID,
			&discount.Qty,
			&discount.Result,
			&discount.Type,
			&discount.ExpiredAt,
		)

		if err != nil {
			if !strings.Contains(err.Error(), "converting NULL") {
				return products, err
			}
		}

		if category.ID != 0 {
			product.Category = &category
		}
		if discount.ID != 0 {
			product.Discount = &discount
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (r *ProductRepository) Count(params model.GetProductQuery) (int, error) {
	var values []interface{}
	var where []string

	query := "SELECT COUNT(*) FROM payments"

	if params.CategoryId != nil {
		values = append(values, params.CategoryId)
		where = append(where, fmt.Sprintf("%s = ?", "category_id"))
	}

	if params.Q != nil {
		values = append(values, params.Q)
		where = append(where, fmt.Sprintf("%s LIKE %s", "name", "'?%'"))
	}

	if len(where) > 0 {
		query += " WHERE "
		query += strings.Join(where, " AND ")
	}

	var count int
	err := r.db.QueryRow(query, values...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProductRepository) GetByID(id int64) (*entity.Product, error) {
	var product entity.Product
	var category entity.Category
	var discount entity.ProductDiscount

	query := `
		SELECT p.id, p.name, p.stock, p.price, p.image, p.created_at, p.updated_at, 
			p.category_id, c.name AS category_name, 
			p.discount_id, d.qty AS discount_qty, d.result AS discount_result, d.type AS discount_type, d.expired_at AS discount_exp_at 
		FROM products p 
		LEFT JOIN categories c ON p.category_id = c.id 
		LEFT JOIN product_discounts d ON p.discount_id = d.id 
		WHERE p.id = ?
	`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Stock,
		&product.Price,
		&product.Image,
		&product.CreatedAt,
		&product.UpdatedAt,

		&category.ID,
		&category.Name,

		&discount.ID,
		&discount.Qty,
		&discount.Result,
		&discount.Type,
		&discount.ExpiredAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.EntityNotFound("Product")
	}

	if err != nil {
		if !strings.Contains(err.Error(), "converting NULL") {
			return nil, err
		}
	}

	if category.ID != 0 {
		product.Category = &category
	}
	if discount.ID != 0 {
		product.Discount = &discount
	}

	return &product, nil
}

func (r *ProductRepository) Create(product entity.Product) (int64, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var discountID *int64
	if product.Discount != nil {
		// Insert discount
		discountResult, err := tx.ExecContext(
			ctx,
			"INSERT INTO product_discounts (qty, type, result, expired_at) VALUES (?, ?, ?, ?)",
			product.Discount.Qty,
			product.Discount.Type,
			product.Discount.Result,
			product.Discount.ExpiredAt,
		)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		id, err := discountResult.LastInsertId()
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		discountID = &id
	}

	// Insert product
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO products (category_id, name, image, price, stock, discount_id) VALUES (?, ?, ?, ?, ?, ?)",
		product.Category.ID,
		product.Name,
		product.Image,
		product.Price,
		product.Stock,
		discountID,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var productId int64
	productId, err = result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, nil
}

func (r *ProductRepository) UpdateByID(product entity.Product) error {
	_, err := r.db.Exec(
		"UPDATE products SET category_id = ?, name = ?, image = ?, price = ?, stock = ? WHERE id = ?",
		product.Category.ID,
		product.Name,
		product.Image,
		product.Price,
		product.Stock,
		product.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteByID(id int64) error {
	if _, err := r.db.Exec(
		"DELETE FROM products WHERE id = ?",
		id,
	); err != nil {
		return err
	}
	return nil
}
