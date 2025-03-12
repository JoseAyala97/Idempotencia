package product

import "time"

type Products struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string, price float64, stock int) (*Products, error) {
	return &Products{
		Name:  name,
		Price: price,
		Stock: stock,
	}, nil
}
