package products

type ProductsRepository interface {
	AddProduct()
}

type Repository struct {
}

var _ ProductsRepository = (*Repository)(nil)

func (s *Repository) AddProduct() {
}
