package repository

import (
	"database/sql"
)

type CashierRepository struct {
	db *sql.DB
}

func NewCashierRepository(db *sql.DB) *CashierRepository {
	return &CashierRepository{db}
}

func (r *CashierRepository) FindAll() {
}

func (r *CashierRepository) GetByID() {
}

func (r *CashierRepository) Create() {
}

func (r *CashierRepository) UpdateByID() {
}

func (r *CashierRepository) DeleteByID() {
}
