package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/services/storage"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
)

type PostApiSendCoinHandler struct {
	AuthService    auth.AuthService
	JWTService     jwt.JWTService
	StorageService storage.StorageService
}

func (h *PostApiSendCoinHandler) PostApiSendCoin(token string, sendCoin entities.SendCoin) error {
	userAuth, err := h.JWTService.ParseToken(token)
	if err != nil {
		return fmt.Errorf("ParseToken fail: %w", err)
	}

	err = h.AuthService.AuthorizeUser(userAuth)
	if err != nil {
		return fmt.Errorf("AuthorizeUser fail: %w", err)
	}

	sendCoin.FromUser = userAuth.Username
	err = h.StorageService.SendCoins(sendCoin)
	if err != nil {
		return fmt.Errorf("SendCoins fail: %w", err)
	}

	return nil
}
