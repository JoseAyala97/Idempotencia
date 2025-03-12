package order

type OrderRepository interface {
	Create(order *Order) (int, error)
	FindById(id int) (*Order, error)
}
