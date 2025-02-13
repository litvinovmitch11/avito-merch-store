package handlers

import (
	"github.com/litvinovmitch11/avito-merch-store/internal/services/products"
)

type PostAdminProductsAddHandler struct {
	ProductsService products.ProductsService
}

func (h *PostAdminProductsAddHandler) PostAdminProductsAdd() {
}
