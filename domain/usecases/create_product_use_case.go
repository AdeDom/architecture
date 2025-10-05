package usecases

import (
	"architecture/data/entities"
	"architecture/data/repositories"
	"errors"
)

type CreateProductUseCase interface {
	Execute(product entities.Product) (*entities.Product, error)
}

type createProductUseCaseImpl struct {
	repository repositories.ProductRepository
}

func NewCreateProductUseCase(repository repositories.ProductRepository) CreateProductUseCase {
	return &createProductUseCaseImpl{repository: repository}
}

func (uc *createProductUseCaseImpl) Execute(product entities.Product) (*entities.Product, error) {
	if product.Name == "" {
		return nil, errors.New("product name is empty")
	}

	if product.Price <= 0 {
		return nil, errors.New("product price is incorrect")
	}

	count, err := uc.repository.GetProductCountByName(product.Name)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("product name already exists")
	}

	result, err := uc.repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
