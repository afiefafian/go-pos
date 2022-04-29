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

func (r *PaymentRepository) FindAll(params *model.PaginationQuery) ([]entity.Payment, error) {
	rows, err := r.db.Query(
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

func (r *PaymentRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM cashiers").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PaymentRepository) GetByID(id int64) (*entity.Payment, error) {
	var payment entity.Payment
	row := r.db.QueryRow("SELECT id, name, type, logo, created_at, updated_at FROM payments WHERE id = ?", id)
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

func (r *PaymentRepository) Create(payment entity.Payment) (int64, error) {
	result, err := r.db.Exec(
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

func (r *PaymentRepository) UpdateByID(payment entity.Payment) error {
	_, err := r.db.Exec(
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

func (r *PaymentRepository) DeleteByID(id int64) error {
	if _, err := r.db.Exec(
		"DELETE FROM payments WHERE id = ?",
		id,
	); err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) FindAllWithRevenues() ([]model.PaymentReportResponse, error) {
	query := `
		SELECT p.id, p.name, p.type, p.logo, COALESCE(SUM(o.total_price), 0) AS total_price
		FROM payments p
		LEFT JOIN orders o ON p.id = o.payment_id
		WHERE total_price > 0 
		GROUP BY p.id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []model.PaymentReportResponse

	for rows.Next() {
		var payment model.PaymentReportResponse
		err := rows.Scan(
			&payment.ID,
			&payment.Name,
			&payment.Type,
			&payment.Logo,
			&payment.TotalAmount,
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
