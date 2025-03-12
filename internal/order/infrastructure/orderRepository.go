package infrastructure

import (
	"IdEmpotencia/internal/order"
	"IdEmpotencia/pkg/database"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(o *order.Order) (int, error) {
	tx := database.DB.Begin()

	if err := tx.Create(o).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return o.ID, nil
}

func (r *OrderRepository) FindById(id int) (*order.Order, error) {
	var o order.Order
	err := database.DB.Preload("OrderItems").First(&o, id).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}
