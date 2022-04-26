package repository

import (
	"database/sql"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/model"
)

type CashierRepository struct {
	db *sql.DB
}

func NewCashierRepository(db *sql.DB) *CashierRepository {
	return &CashierRepository{db}
}

func (r *CashierRepository) FindAll(params *model.PaginationQuery) ([]entity.Cashier, error) {
	rows, err := r.db.Query(
		"SELECT id, name, passcode, created_at, updated_at FROM cashiers LIMIT ? OFFSET ?",
		params.Limit,
		params.Skip,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cashiers []entity.Cashier

	for rows.Next() {
		var cashier entity.Cashier
		err := rows.Scan(
			&cashier.ID,
			&cashier.Name,
			&cashier.Passcode,
			&cashier.CreatedAt,
			&cashier.UpdatedAt,
		)

		if err != nil {
			return cashiers, err
		}

		cashiers = append(cashiers, cashier)
	}

	if err = rows.Err(); err != nil {
		return cashiers, err
	}

	return cashiers, nil
}

func (r *CashierRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM cashiers").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CashierRepository) GetByID(id int64) (*entity.Cashier, error) {
	var cashier entity.Cashier
	row := r.db.QueryRow("SELECT id, name, passcode, created_at, updated_at FROM cashiers WHERE id = ?", id)
	err := row.Scan(
		&cashier.ID,
		&cashier.Name,
		&cashier.Passcode,
		&cashier.CreatedAt,
		&cashier.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.EntityNotFound("Cashier")
	}

	if err != nil {
		return nil, err
	}

	return &cashier, nil
}

func (r *CashierRepository) Create(cashier entity.Cashier) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO cashiers (name, passcode) VALUES (?, ?)",
		cashier.Name,
		cashier.Passcode,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *CashierRepository) UpdateByID(cashier entity.Cashier) error {
	_, err := r.db.Exec(
		"UPDATE cashiers SET name = ? WHERE id = ?",
		cashier.Name,
		cashier.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CashierRepository) DeleteByID(id int64) error {
	if _, err := r.db.Exec(
		"DELETE FROM cashiers WHERE id = ?",
		id,
	); err != nil {
		return err
	}
	return nil
}
