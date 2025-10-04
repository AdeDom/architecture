package core

import (
	"errors"
)

// Primary port
type ProductService interface {
	CreateProduct(product Product) (*Product, error)
	GetProductById(id int) (*Product, error)
	GetProducts(name string, price float64) ([]Product, error)
}

type productServiceImpl struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &productServiceImpl{repository: repository}
}

func (s *productServiceImpl) CreateProduct(product Product) (*Product, error) {
	if product.Name == "" {
		return nil, errors.New("product name is empty")
	}

	if product.Price <= 0 {
		return nil, errors.New("product price is incorrect")
	}

	count, err := s.repository.GetProductCountByName(product.Name)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("product name already exists")
	}

	result, err := s.repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productServiceImpl) GetProductById(id int) (*Product, error) {
	return s.repository.GetProductById(id)
}

func (s *productServiceImpl) GetProducts(name string, price float64) ([]Product, error) {
	return s.repository.GetProducts(name, price)
}
