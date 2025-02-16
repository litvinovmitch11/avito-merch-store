package handlers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/handlers"
	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"
	mock_auth "github.com/litvinovmitch11/avito-merch-store/mocks/services/auth"
	mock_jwt "github.com/litvinovmitch11/avito-merch-store/mocks/services/jwt"
)

func TestPostApiAuth_AuthorizeTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock_auth.NewMockAuthService(ctrl)
	JWTService := mock_jwt.NewMockJWTService(ctrl)

	postPostApiAuthHandler := handlers.PostApiAuthHandler{
		AuthService: authService,
		JWTService:  JWTService,
	}

	userAuth := entities.UserAuth{
		Username: "Username",
		Password: "Password",
	}

	tokenStr := "token"
	expectedToken := entities.UserToken{
		Token: &tokenStr,
	}

	gomock.InOrder(
		authService.EXPECT().AuthorizeUser(userAuth).Return("", nil),
		JWTService.EXPECT().NewToken(userAuth).Return(tokenStr, nil),
	)

	token, err := postPostApiAuthHandler.PostApiAuth(userAuth)
	if err != nil {
		t.Errorf("PostApiAuth fail: %v", err)
	}

	if token.Token == nil {
		t.Errorf("nil token")
	}

	if *expectedToken.Token != *token.Token {
		t.Errorf("expected: %s, real: %s", *expectedToken.Token, *token.Token)
	}
}

func TestPostApiAuth_CreateTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock_auth.NewMockAuthService(ctrl)
	JWTService := mock_jwt.NewMockJWTService(ctrl)

	postPostApiAuthHandler := handlers.PostApiAuthHandler{
		AuthService: authService,
		JWTService:  JWTService,
	}

	userAuth := entities.UserAuth{
		Username: "Username",
		Password: "Password",
	}

	tokenStr := "token"
	expectedToken := entities.UserToken{
		Token: &tokenStr,
	}

	gomock.InOrder(
		authService.EXPECT().AuthorizeUser(userAuth).Return("", authrepo.ErrUserNotFound),
		authService.EXPECT().CreateUser(userAuth).Return("", nil),
		JWTService.EXPECT().NewToken(userAuth).Return(tokenStr, nil),
	)

	token, err := postPostApiAuthHandler.PostApiAuth(userAuth)
	if err != nil {
		t.Errorf("PostApiAuth fail: %v", err)
	}

	if token.Token == nil {
		t.Errorf("nil token")
	}

	if *expectedToken.Token != *token.Token {
		t.Errorf("expected: %s, real: %s", *expectedToken.Token, *token.Token)
	}
}
