package handlers

import (
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
)

type PostApiAuthHandler struct {
	AuthService auth.AuthService
}

func (h *PostApiAuthHandler) PostApiAuth(userAuth entities.UserAuth) (entities.UserToken, error) {
	token := h.AuthService.GetToken(userAuth)

	return token, nil
}
