package service

import (
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
					var normalPriceQty int32
					var discountedQty int32
					if qty >= product.Discount.Qty {
						normalPriceQty = qty - product.Discount.Qty
						discountedQty = product.Discount.Qty
					} else if qty < product.Discount.Qty {
						normalPriceQty = 0
						discountedQty = qty
					}

					totalFinalPrice = (int64(normalPriceQty) * product.Price) + (int64(discountedQty) * product.Discount.Result)
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
