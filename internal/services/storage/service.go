package storage

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
	storagerepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/storage"
)

type StorageService interface {
	SendCoins(sendCoin entities.SendCoin) error
	BuyMerch(userID, merchID string) error
}

type Service struct {
	ProductsRepository productsrepo.ProductsRepository
	StorageRepository  storagerepo.StorageRepository
}

var _ StorageService = (*Service)(nil)

func (s *Service) SendCoins(sendCoin entities.SendCoin) error {
	// user, err := s.AuthRepository.GetUserByUsername(userAuth.Username)
	// if err != nil {
	// 	return fmt.Errorf("GetUserByUsername fail: %w", err)
	// }

	// user, err := s.AuthRepository.GetUserByUsername(userAuth.Username)
	// if err != nil {
	// 	return fmt.Errorf("GetUserByUsername fail: %w", err)
	// }

	err := s.StorageRepository.SendCoins(sendCoin)
	if err != nil {
		return fmt.Errorf("SendCoins fail: %w", err)
	}

	return nil
}

func (s *Service) BuyMerch(userID, merchID string) error {
	// s.ProductsRepository.GetProduct()

	// entity := entities.SendCoin{
	// 	FromUser: userID,
	// 	ToUser: "",
	// 	Amount: ,
	// }

	// err := s.StorageRepository.SendCoins()
	// if err != nil {
	// 	return fmt.Errorf("SendCoins fail: %w", err)
	// }

	return nil
}
