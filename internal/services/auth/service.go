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
	AuthorizeUser(entities.UserAuth) error
}

type Service struct {
	AuthRepository authrepo.AuthRepository
}

var _ AuthService = (*Service)(nil)

func (s *Service) CreateUser(userAuth entities.UserAuth) (string, error) {
	password := []byte(userAuth.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("GenerateFromPassword fail: %w", err)
	}

	user := entities.User{
		ID:       uuid.NewString(),
		Username: userAuth.Username,
	}

	pd := entities.UserPersonalData{
		ID:             uuid.NewString(),
		UserID:         user.ID,
		HashedPassword: string(hashedPassword),
	}

	balance := entities.Balance{
		ID:     uuid.NewString(),
		UserID: user.ID,
		Amount: 1000,
	}

	err = s.AuthRepository.CreateUser(user, pd, balance)
	if err != nil {
		return "", fmt.Errorf("CreateUser fail: %w", err)
	}

	return user.ID, nil
}

func (s *Service) AuthorizeUser(userAuth entities.UserAuth) error {
	user, err := s.AuthRepository.GetUserByUsername(userAuth.Username)
	if err != nil {
		return fmt.Errorf("GetUserByUsername fail: %w", err)
	}

	personalData, err := s.AuthRepository.GetPersonalData(user.ID)
	if err != nil {
		return fmt.Errorf("GetPersonalData fail: %w", err)
	}

	password := []byte(userAuth.Password)
	hashedPassword := []byte(personalData.HashedPassword)
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("CompareHashAndPassword fail: %w", err)
	}

	return nil
}
