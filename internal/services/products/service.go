package products

import (
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
)

type ProductsService interface {
	AddProduct()
}

type Service struct {
	ProductsRepository productsrepo.ProductsRepository
}

var _ ProductsService = (*Service)(nil)

func (s *Service) AddProduct() {
	s.ProductsRepository.AddProduct()
}
