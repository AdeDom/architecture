package adapters

import (
	"architecture/core/product"
	"gorm.io/gorm"
)

// Secondary adapter
type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) core.ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) CreateProduct(product core.Product) (*core.Product, error) {
	if result := r.db.Create(&product); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepositoryImpl) GetProductById(id int) (*core.Product, error) {
	var product core.Product
	if result := r.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepositoryImpl) GetProductCountByName(name string) (int64, error) {
	var count int64
	result := r.db.Model(&core.Product{}).Where("name = ?", name).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (r *productRepositoryImpl) GetProducts(name string, price float64) ([]core.Product, error) {
	var products []core.Product
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
