package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/afiefafian/go-pos/src/entity"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
)

type OrderService struct {
	OrderRepository   *repository.OrderRepository
	ProductRepository *repository.ProductRepository
}

func NewOrderService(
	orderRepository *repository.OrderRepository,
	paymentRepository *repository.ProductRepository,
) *OrderService {
	return &OrderService{
		OrderRepository:   orderRepository,
		ProductRepository: paymentRepository,
	}
}

// func (s *OrderService) GetByID(id int64) (*model.GetOrderPaymentResponse, error) {
// 	order, err := s.OrderRepository.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// }

func (s *OrderService) IsInvoiceDownloaded(id int64) (bool, error) {
	order, err := s.OrderRepository.GetByID(id)
	if err != nil {
		return false, err
	}
	return order.IsDownloaded, nil
}

func (s *OrderService) ChangeInvoiceDownloadStatus(id int64) error {
	status, err := s.IsInvoiceDownloaded(id)
	if err != nil {
		return err
	}
	if status {
		return nil
	}
	return s.OrderRepository.ChangeDownloadStatusByID(id)
}

func (s *OrderService) CheckSubTotal(request model.CreateSubTotalRequest) (*model.CreateSubTotalResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	// Collect all product id from request
	var productIDs []int64
	productQtyMap := make(map[int64]int32)
	for _, v := range request.Products {
		if v.Qty > 0 {
			productIDs = append(productIDs, v.ProductID)
			productQtyMap[v.ProductID] = v.Qty
		}
	}

	// Get products
	products, err := s.ProductRepository.FindByIDS(productIDs)
	if err != nil {
		return nil, err
	}

	// If product from request and db is not same
	// then return error
	if len(products) != len(productIDs) {
		return nil, errors.New("Invalid Product ID")
	}

	var subTotal int64 = 0
	var productsResponse = make([]model.CreateOrderProductResponse, 0)
	for _, product := range products {
		qty := productQtyMap[product.ID]
		finalPrice := product.FinalPrice()
		totalNormalPrice := product.Price * int64(qty)
		totalFinalPrice := finalPrice * int64(qty)

		productResponse := model.CreateOrderProductResponse{
			ID:               product.ID,
			Name:             product.Name,
			Stock:            product.Stock,
			Price:            product.Price,
			Qty:              qty,
			TotalNormalPrice: totalNormalPrice,
			TotalFinalPrice:  totalFinalPrice,
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

			if product.Discount.IsValid() {
				switch product.Discount.Type {
				case "BUY_N":
					if qty == product.Discount.Qty {
						totalFinalPrice = product.Discount.Result
					}
				case "PERCENT":
					totalFinalPrice = int64(qty) * finalPrice
				}
			}
		}

		subTotal += totalFinalPrice

		productsResponse = append(productsResponse, productResponse)
	}

	response := &model.CreateSubTotalResponse{
		Subtotal: subTotal,
		Products: productsResponse,
	}

	return response, nil
}

func (s *OrderService) CreateOrder(request model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	if err := helper.ValidateStruct(request); err != nil {
		return nil, err
	}

	// Get product info from sub total
	subtotalRes, err := s.CheckSubTotal(model.CreateSubTotalRequest{
		Products: request.Products,
	})
	if err != nil {
		return nil, err
	}

	if request.TotalPaid < subtotalRes.Subtotal {
		return nil, errors.New("Total paid amount is smaller than total price")
	}

	// Collect required order data
	order := entity.Order{
		CashierID:     request.CashierID,
		PaymentTypeID: request.PaymentID,
		TotalPrice:    0,
		TotalPaid:     request.TotalPaid,
		TotalReturn:   0,
		ReceiptID:     "",
		IsDownloaded:  false,
	}
	// Generate Receipt ID
	order.ReceiptID = order.GenerateReceiptID()
	order.TotalPrice = subtotalRes.Subtotal
	order.TotalReturn = order.TotalPaid - subtotalRes.Subtotal

	// Collect required order products data
	var orderProducts []entity.OrderProduct
	for _, v := range subtotalRes.Products {
		// Check stock qty
		if v.Qty > v.Stock {
			return nil, errors.New(fmt.Sprintf("Insufficient stock: Order qty total (%d) is bigger tham product stock (%d)", v.Qty, v.Stock))
		}

		product := entity.OrderProduct{
			ProductID:        v.ID,
			Qty:              v.Qty,
			Price:            v.Price,
			TotalNormalPrice: v.TotalNormalPrice,
			TotalFinalPrice:  v.TotalFinalPrice,
		}

		if v.Discount != nil {
			product.DiscountID = v.Discount.ID
		}

		orderProducts = append(orderProducts, product)
	}

	// Store to database
	orderID, err := s.OrderRepository.Create(&order, orderProducts)
	if err != nil {
		return nil, err
	}

	// Response
	currentTime := time.Now()
	orderResponse := model.GetOrderResponse{
		ID:            orderID,
		CashierID:     request.CashierID,
		PaymentTypeID: request.PaymentID,
		TotalPrice:    order.TotalPaid,
		TotalPaid:     request.TotalPaid,
		TotalReturn:   order.TotalReturn,
		ReceiptID:     order.ReceiptID,
		CreatedAt:     currentTime,
		UpdatedAt:     currentTime,
	}

	productsResponse := make([]model.CreateOrderProductResponse, 0)
	for _, v := range subtotalRes.Products {
		product := model.CreateOrderProductResponse{
			ID:               v.ID,
			Name:             v.Name,
			Price:            v.Price,
			Discount:         v.Discount,
			Qty:              v.Qty,
			TotalNormalPrice: v.TotalNormalPrice,
			TotalFinalPrice:  v.TotalFinalPrice,
		}
		productsResponse = append(productsResponse, product)
	}

	response := model.CreateOrderResponse{
		Order:    orderResponse,
		Products: productsResponse,
	}
	return &response, nil
}
