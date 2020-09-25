package product

type Product struct {
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByCategory(categoryName string) ([]Product, error)
	GetByType(categoryName string) ([]Product, error)
	GetByName(productName string) (*Product, error)
}
