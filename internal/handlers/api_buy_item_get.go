package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/storage"
)

type GetApiBuyItemHandler struct {
	AuthService    auth.AuthService
	JWTService     jwt.JWTService
	StorageService storage.StorageService
}

func (h *GetApiBuyItemHandler) GetApiBuyItem(token string, item string) error {
	userAuth, err := h.JWTService.ParseToken(token)
	if err != nil {
		return fmt.Errorf("ParseToken fail: %w", err)
	}

	err = h.AuthService.AuthorizeUser(userAuth)
	if err != nil {
		return fmt.Errorf("AuthorizeUser fail: %w", err)
	}

	err = h.StorageService.BuyMerch(userAuth.Username, item)
	if err != nil {
		return fmt.Errorf("BuyMerch fail: %w", err)
	}

	return nil
}
