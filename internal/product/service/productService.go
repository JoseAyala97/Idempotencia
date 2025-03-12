package service

import (
	"IdEmpotencia/internal/product"
	"errors"
)

type ProductService struct {
	repo product.ProductRepository
}

func NewProductService(repo product.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]*product.Products, error) {
	return s.repo.GetAll()
}

func (s *ProductService) UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return s.repo.UpdateStock(id, newStock)
}
