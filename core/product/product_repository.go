package core

// Secondary port
type ProductRepository interface {
	CreateProduct(product Product) (*Product, error)
	GetProductById(id int) (*Product, error)
	GetProductCountByName(name string) (int64, error)
	GetProducts(name string, price float64) ([]Product, error)
}
