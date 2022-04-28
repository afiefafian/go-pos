package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/exception"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) GetByID(id int64) (*entity.Order, error) {
	var order entity.Order
	row := r.db.QueryRow("SELECT id, cashier_id, payment_id, total_price, total_paid, total_return, receipt_id, is_downloaded, created_at, updated_at FROM orders WHERE id = ?", id)
	err := row.Scan(
		&order.ID,
		&order.CashierID,
		&order.PaymentTypeID,
		&order.TotalPrice,
		&order.TotalPaid,
		&order.TotalReturn,
		&order.ReceiptID,
		&order.IsDownloaded,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.EntityNotFound("Order")
	}

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) Create(order *entity.Order, products []entity.OrderProduct) (int64, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	// Insert order
	orderResult, err := tx.ExecContext(
		ctx,
		"INSERT INTO orders (cashier_id, payment_id, total_price, total_paid, total_return, receipt_id) VALUES (?, ?, ?, ?, ?, ?)",
		order.CashierID,
		order.PaymentTypeID,
		order.TotalPrice,
		order.TotalPaid,
		order.TotalPrice,
		order.ReceiptID,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var orderID int64
	orderID, err = orderResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Insert many order products data
	var (
		orderProductsInsertPlaceholder []string
		orderProductsInsertVals        []interface{}
		productIDs                     []interface{}
	)
	productQtyMap := make(map[int64]int32)

	for _, product := range products {
		orderProductsInsertPlaceholder = append(orderProductsInsertPlaceholder, "(?, ?, ?, ?, ?, ?, ?)")
		orderProductsInsertVals = append(
			orderProductsInsertVals,
			orderID,
			product.ProductID,
			product.DiscountID,
			product.Qty,
			product.Price,
			product.TotalNormalPrice,
			product.TotalFinalPrice,
		)
		productIDs = append(productIDs, product.ProductID)
		productQtyMap[product.ProductID] = product.Qty
	}
	insertOrderProductsQuery := fmt.Sprintf(
		"INSERT INTO order_products (order_id, product_id, discount_id, qty, price, total_normal_price, total_final_price) VALUES %s",
		strings.Join(orderProductsInsertPlaceholder, ", "),
	)
	if _, err := tx.ExecContext(ctx, insertOrderProductsQuery, orderProductsInsertVals...); err != nil {
		tx.Rollback()
		return 0, err
	}

	// Update products stocks
	var (
		productsUpdatePlaceholder []string
		productsUpdateVals        []interface{}
	)
	for productID, qty := range productQtyMap {
		productsUpdatePlaceholder = append(productsUpdatePlaceholder, "WHEN ? THEN stock - ?")
		productsUpdateVals = append(productsUpdateVals, productID, qty)
	}
	productsUpdateVals = append(productsUpdateVals, productIDs...)
	updateProductsStockQuery := fmt.Sprintf(
		"UPDATE products SET stock = CASE id %s END WHERE id IN (?%s) AND stock > 0",
		strings.Join(productsUpdatePlaceholder, " "),
		strings.Repeat(", ?", len(productIDs)-1),
	)
	if _, err := tx.ExecContext(ctx, updateProductsStockQuery, productsUpdateVals...); err != nil {
		tx.Rollback()
		return 0, err
	}

	// Commit changes
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return orderID, nil
}

func (r *OrderRepository) ChangeDownloadStatusByID(id int64) error {
	if _, err := r.db.Exec("UPDATE orders SET is_downloaded = TRUE WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
