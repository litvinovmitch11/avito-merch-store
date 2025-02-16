package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/storage"
)

type GetApiInfoHandler struct {
	AuthService    auth.AuthService
	JWTService     jwt.JWTService
	StorageService storage.StorageService
}

func (h *GetApiInfoHandler) GetApiInfo(token string) (entities.Info, error) {
	userAuth, err := h.JWTService.ParseToken(token)
	if err != nil {
		return entities.Info{}, fmt.Errorf("ParseToken fail: %w", err)
	}

	userID, err := h.AuthService.AuthorizeUser(userAuth)
	if err != nil {
		return entities.Info{}, fmt.Errorf("AuthorizeUser fail: %w", err)
	}

	balance, err := h.StorageService.GetBalance(userID)
	if err != nil {
		return entities.Info{}, fmt.Errorf("GetBalance fail: %w", err)
	}

	inventory, err := h.StorageService.GetInventory(userID)
	if err != nil {
		return entities.Info{}, fmt.Errorf("GetInventory fail: %w", err)
	}

	history, err := h.StorageService.GetHistory(userID)
	if err != nil {
		return entities.Info{}, fmt.Errorf("GetHistory fail: %w", err)
	}

	return entities.Info{
		Coins:       balance,
		Inventory:   inventory,
		CoinHistory: history,
	}, nil
}
