package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type JWTService interface {
	NewToken(userAuth entities.UserAuth) (string, error)
	ParseToken(tokenString string) (entities.UserAuth, error)
}

type Service struct{}

var _ JWTService = (*Service)(nil)

var secret = os.Getenv("JWTSECRET")

func (s *Service) NewToken(userAuth entities.UserAuth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userAuth.Username,
		"password": userAuth.Password,
		"exp":      time.Now().Add(time.Hour).UTC().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("SignedString fail: %w", err)
	}

	return tokenString, nil
}

func (s *Service) ParseToken(tokenString string) (entities.UserAuth, error) {
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return entities.UserAuth{}, fmt.Errorf("jwt.Parse fail: %w: %w", err, ErrInvalidToken)
	}

	if ve, ok := err.(*jwt.ValidationError); !tokenParsed.Valid && ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return entities.UserAuth{}, fmt.Errorf("that's not even a token: %w", ErrInvalidToken)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return entities.UserAuth{}, fmt.Errorf("timing is everything: %w", ErrInvalidToken)
		}
		return entities.UserAuth{}, fmt.Errorf("couldn't handle this token: %w: %w", err, ErrInvalidToken)
	} else if !tokenParsed.Valid {
		return entities.UserAuth{}, fmt.Errorf("couldn't handle this token: %w: %w", err, ErrInvalidToken)
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return entities.UserAuth{}, fmt.Errorf("couldn't parse this token: %w", ErrInvalidToken)
	}

	username, ok := claims["username"].(string)
	if !ok {
		return entities.UserAuth{}, fmt.Errorf("couldn't parse this token: %w", ErrInvalidToken)
	}

	password, ok := claims["password"].(string)
	if !ok {
		return entities.UserAuth{}, fmt.Errorf("couldn't parse this token: %w", ErrInvalidToken)
	}

	return entities.UserAuth{
		Username: username,
		Password: password,
	}, nil
}
