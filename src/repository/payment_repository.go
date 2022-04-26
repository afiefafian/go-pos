package repository

import (
	"database/sql"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/model"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (c *PaymentRepository) FindAll(params *model.PaginationQuery) ([]entity.Payment, error) {
	rows, err := c.db.Query(
		"SELECT id, name, type, logo, created_at, updated_at FROM payments LIMIT ? OFFSET ?",
		params.Limit,
		params.Skip,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []entity.Payment

	for rows.Next() {
		var payment entity.Payment
		err := rows.Scan(
			&payment.ID,
			&payment.Name,
			&payment.Type,
			&payment.Logo,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)

		if err != nil {
			return payments, err
		}

		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		return payments, err
	}

	return payments, nil
}

func (c *PaymentRepository) Count() (int, error) {
	var count int
	err := c.db.QueryRow("SELECT COUNT(*) FROM cashiers").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *PaymentRepository) GetByID(id int64) (*entity.Payment, error) {
	var payment entity.Payment
	row := c.db.QueryRow("SELECT id, name, type, logo, created_at, updated_at FROM payments WHERE id = ?", id)
	err := row.Scan(
		&payment.ID,
		&payment.Name,
		&payment.Type,
		&payment.Logo,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.EntityNotFound("Payment")
	}

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (c *PaymentRepository) Create(payment entity.Payment) (int64, error) {
	result, err := c.db.Exec(
		"INSERT INTO payments (name, type, logo) VALUES (?, ?, ?)",
		payment.Name,
		payment.Type,
		payment.Logo,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (c *PaymentRepository) UpdateByID(payment entity.Payment) error {
	_, err := c.db.Exec(
		"UPDATE payments SET name = ?, type = ?, logo = ? WHERE id = ?",
		payment.Name,
		payment.Type,
		payment.Logo,
		payment.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *PaymentRepository) DeleteByID(id int64) error {
	if _, err := c.db.Exec(
		"DELETE FROM payments WHERE id = ?",
		id,
	); err != nil {
		return err
	}
	return nil
}
