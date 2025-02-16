package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"

	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/products"
)

type PostAdminProductsAddHandler struct {
	AuthService     auth.AuthService
	JWTService      jwt.JWTService
	ProductsService products.ProductsService
}

func (h *PostAdminProductsAddHandler) PostAdminProductsAdd(token string, product entities.Product) (string, error) {
	userAuth, err := h.JWTService.ParseToken(token)
	if err != nil {
		return "", fmt.Errorf("ParseToken fail: %w", err)
	}

	_, err = h.AuthService.AuthorizeUser(userAuth)
	if err != nil {
		return "", fmt.Errorf("AuthorizeUser fail: %w", err)
	}

	id, err := h.ProductsService.AddProduct(product)
	if err != nil {
		return id, fmt.Errorf("AddProduct fail: %w", err)
	}

	return id, nil
}
