package repositories

import (
	"architecture/data/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product entities.Product) (*entities.Product, error)
	GetProductById(id int) (*entities.Product, error)
	GetProductCountByName(name string) (int64, error)
	GetProducts(name string, price float64) ([]entities.Product, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) CreateProduct(product entities.Product) (*entities.Product, error) {
	if result := r.db.Create(&product); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepositoryImpl) GetProductById(id int) (*entities.Product, error) {
	var product entities.Product
	if result := r.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepositoryImpl) GetProductCountByName(name string) (int64, error) {
	var count int64
	result := r.db.Model(&entities.Product{}).Where("name = ?", name).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (r *productRepositoryImpl) GetProducts(name string, price float64) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if price > 0 {
		query = query.Where("price = ?", price)
	}

	if result := query.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
