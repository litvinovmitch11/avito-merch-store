package products

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
)

type ProductsService interface {
	AddProduct(product entities.Product) (string, error)
}

type Service struct {
	ProductsRepository productsrepo.ProductsRepository
}

var _ ProductsService = (*Service)(nil)

func (s *Service) AddProduct(product entities.Product) (string, error) {
	product.Id = uuid.NewString()

	err := s.ProductsRepository.AddProduct(product)
	if err != nil {
		return "", fmt.Errorf("AddProduct fail: %w", err)
	}

	return product.Id, nil
}
