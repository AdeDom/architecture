package usecases

import (
	"architecture/data/entities"
	"architecture/data/repositories"
)

type GetProductsUseCase interface {
	Execute(name string, price float64) ([]entities.Product, error)
}

type getProductsUseCaseImpl struct {
	repository repositories.ProductRepository
}

func NewGetProductsUseCase(repository repositories.ProductRepository) GetProductsUseCase {
	return &getProductsUseCaseImpl{repository: repository}
}

func (uc *getProductsUseCaseImpl) Execute(name string, price float64) ([]entities.Product, error) {
	return uc.repository.GetProducts(name, price)
}
