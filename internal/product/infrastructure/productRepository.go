package infrastructure

import (
	"IdEmpotencia/internal/product"
	"IdEmpotencia/pkg/database"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetAll() ([]*product.Products, error) {
	var products []*product.Products
	result := database.DB.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) GetByID(id int) (*product.Products, error) {
	var p product.Products
	if err := database.DB.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) UpdateStock(id int, stock int) error {
	return database.DB.Model(&product.Products{}).Where("id = ?", id).Update("stock", stock).Error
}
