package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
)

type JWTService interface {
	NewToken(userAuth entities.UserAuth) (string, error)
	ParseToken(tokenString string) (entities.UserAuth, error)
}

type Service struct{}

var _ JWTService = (*Service)(nil)

var secret = "uuid.NewString()"

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
		return entities.UserAuth{}, err
	}

	if tokenParsed.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
		return entities.UserAuth{}, errors.New("123")
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return entities.UserAuth{}, errors.New("321")
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return entities.UserAuth{}, errors.New("qwe")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return entities.UserAuth{}, errors.New("qwe")
	}

	password, ok := claims["password"].(string)
	if !ok {
		return entities.UserAuth{}, errors.New("qwe")
	}

	return entities.UserAuth{
		Username: username,
		Password: password,
	}, nil
}
