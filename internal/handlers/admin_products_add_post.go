package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/products"
)

type PostAdminProductsAddHandler struct {
	ProductsService products.ProductsService
}

func (h *PostAdminProductsAddHandler) PostAdminProductsAdd(product entities.Product) (string, error) {
	id, err := h.ProductsService.AddProduct(product)
	if err != nil {
		return id, fmt.Errorf("AddProduct fail: %w", err)
	}

	return id, nil
}
