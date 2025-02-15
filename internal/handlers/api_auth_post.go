package handlers

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
)

type PostApiAuthHandler struct {
	AuthService auth.AuthService
}

func (h *PostApiAuthHandler) PostApiAuth(userAuth entities.UserAuth) (entities.UserToken, error) {
	// need add check "person was created"

	userID, err := h.AuthService.CreateUser(userAuth)
	if err != nil {
		return entities.UserToken{}, fmt.Errorf("CreateUser fail: %w", err)
	}

	return entities.UserToken{Token: &userID}, nil
}
