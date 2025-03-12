package product

type ProductRepository interface {
	GetAll() ([]*Products, error)
	GetByID(id int) (*Products, error)
	UpdateStock(id int, stock int) error
}
