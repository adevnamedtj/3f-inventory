package product

type Product struct {

}

interface ProductRepository {
	func GetAll()([]Product, error)
	func GetByCategory(categoryName string) ([]Product, error)
	func GetByType(categoryName string) ([]Product, error)
	func GetByName(productName string) (*Product, error)
	}