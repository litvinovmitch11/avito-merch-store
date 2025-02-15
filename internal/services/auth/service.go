package auth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(entities.UserAuth) (string, error)
}

type Service struct {
	AuthRepository authrepo.AuthRepository
}

var _ AuthService = (*Service)(nil)

func (s *Service) CreateUser(user entities.UserAuth) (string, error) {
	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("GenerateFromPassword fail: %w", err)
	}

	user.ID = uuid.NewString()

	pd := entities.UserPersonalData{
		ID:             uuid.NewString(),
		UserID:         user.ID,
		HashedPassword: string(hashedPassword),
	}

	err = s.AuthRepository.CreateUser(user, pd)
	if err != nil {
		return "", fmt.Errorf("CreateUser fail: %w", err)
	}

	return "success_token", nil
}

func (s *Service) AuthorizeUser(user entities.UserAuth) (string, error) {
	password := []byte(user.Password)

	hashedPassword := []byte("getpass")

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return "0", fmt.Errorf("CompareHashAndPassword fail: %w", err)
	}

	return "1", nil
}
