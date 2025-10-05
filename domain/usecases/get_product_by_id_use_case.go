package usecases

import (
	"architecture/data/entities"
	"architecture/data/repositories"
)

type GetProductByIdUseCase interface {
	Execute(id int) (*entities.Product, error)
}

type getProductByIdUseCaseImpl struct {
	repository repositories.ProductRepository
}

func NewGetProductByIdUseCase(repository repositories.ProductRepository) GetProductByIdUseCase {
	return &getProductByIdUseCaseImpl{repository: repository}
}

func (uc *getProductByIdUseCaseImpl) Execute(id int) (*entities.Product, error) {
	return uc.repository.GetProductById(id)
}
