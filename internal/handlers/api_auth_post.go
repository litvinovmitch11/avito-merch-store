package handlers

import (
	"errors"
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"

	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"
)

type PostApiAuthHandler struct {
	AuthService auth.AuthService
	JWTService  jwt.JWTService
}

func (h *PostApiAuthHandler) PostApiAuth(userAuth entities.UserAuth) (entities.UserToken, error) {
	err := h.AuthService.AuthorizeUser(userAuth)
	if errors.Is(err, authrepo.ErrUserNotFound) || errors.Is(err, authrepo.ErrPDNotFound) {
		_, err = h.AuthService.CreateUser(userAuth)
		if err != nil {
			return entities.UserToken{}, fmt.Errorf("CreateUser fail: %w", err)
		}
	} else if err != nil {
		return entities.UserToken{}, err
	}

	token, err := h.JWTService.NewToken(userAuth)
	if err != nil {
		return entities.UserToken{}, fmt.Errorf("NewToken fail: %w", err)
	}

	return entities.UserToken{Token: &token}, nil
}
