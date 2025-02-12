package auth

import "github.com/litvinovmitch11/avito-merch-store/internal/entities"

type AuthService interface {
	GetToken(entities.UserAuth) entities.UserToken
}

type Service struct {
}

var _ AuthService = (*Service)(nil)

func (s Service) GetToken(user entities.UserAuth) entities.UserToken {
	token := "token: "
	token += user.Username
	token += " "
	token += user.Password

	return entities.UserToken{
		Token: &token,
	}
}
