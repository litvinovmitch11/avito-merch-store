package Products

type ProductsService interface {
	AddProduct()
}

type Service struct {
}

var _ ProductsService = (*Service)(nil)

func (s Service) AddProduct() {

}
