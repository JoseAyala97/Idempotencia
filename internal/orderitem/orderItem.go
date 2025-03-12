package orderitem

import "time"

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
