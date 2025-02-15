package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
)

type GetApiInfoHandler struct {
	AuthService auth.AuthService
	JWTService  jwt.JWTService
}

func (h *GetApiInfoHandler) GetApiInfo(token string) (string, error) {
	userAuth, err := h.JWTService.ParseToken(token)
	if err != nil {
		return "", fmt.Errorf("ParseToken fail: %w", err)
	}

	err = h.AuthService.AuthorizeUser(userAuth)
	if err != nil {
		return "", fmt.Errorf("AuthorizeUser fail: %w", err)
	}

	// id, err := h.ProductsService.AddProduct(product)
	// if err != nil {
	// 	return id, fmt.Errorf("AddProduct fail: %w", err)
	// }

	return "id", nil
}
