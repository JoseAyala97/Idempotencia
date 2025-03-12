package order

import (
	"time"

	"IdEmpotencia/internal/orderitem"
)

type Order struct {
	ID           int                   `json:"id"`
	CustomerName string                `json:"customer_name"`
	TotalAmount  float64               `json:"total_amount"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	OrderItems   []orderitem.OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
