package service

import (
	"IdEmpotencia/internal/order"
	"IdEmpotencia/internal/product/infrastructure"
	"IdEmpotencia/pkg/database"
	"context"
	"encoding/json"
	"errors"
	"time"
)

type OrderService struct {
	repo        order.OrderRepository
	productRepo *infrastructure.ProductRepository
}

func NewOrderService(repo order.OrderRepository, productRepo *infrastructure.ProductRepository) *OrderService {
	return &OrderService{repo: repo, productRepo: productRepo}
}

func (s *OrderService) CreateOrder(ctx context.Context, idempotencyKey string, o *order.Order) (int, error) {
	data, err := database.RedisClient.Get(ctx, idempotencyKey).Result()
	if err == nil {
		var storedData struct {
			Status  string `json:"status"`
			OrderID int    `json:"order_id"`
		}
		json.Unmarshal([]byte(data), &storedData)

		if storedData.Status == "COMPLETED" {
			return storedData.OrderID, nil
		}
		return 0, errors.New("solicitud en progreso")
	}

	inProgressData := map[string]interface{}{
		"status": "IN_PROGRESS",
	}
	jsonData, _ := json.Marshal(inProgressData)
	database.RedisClient.Set(ctx, idempotencyKey, jsonData, 24*time.Hour)

	var total float64
	for i, item := range o.OrderItems {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil || product.Stock < item.Quantity {
			return 0, errors.New("stock insuficiente")
		}

		o.OrderItems[i].Subtotal = product.Price * float64(item.Quantity)
		total += o.OrderItems[i].Subtotal

		product.Stock -= item.Quantity
		s.productRepo.UpdateStock(product.ID, product.Stock)
	}
	o.TotalAmount = total

	orderID, err := s.repo.Create(o)
	if err != nil {
		return 0, err
	}

	finalData := map[string]interface{}{
		"status":   "COMPLETED",
		"order_id": orderID,
	}
	jsonData, _ = json.Marshal(finalData)
	database.RedisClient.Set(ctx, idempotencyKey, jsonData, 24*time.Hour)

	return orderID, nil
}

func (s *OrderService) GetOrderByID(id int) (*order.Order, error) {
	o, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("pedido no encontrado")
	}
	return o, nil
}
