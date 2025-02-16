package storage

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
	storagerepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/storage"
)

type StorageService interface {
	SendCoins(sendCoin entities.SendCoin) error
	BuyMerch(userID, merchID string) error
}

type Service struct {
	AuthRepository     authrepo.AuthRepository
	ProductsRepository productsrepo.ProductsRepository
	StorageRepository  storagerepo.StorageRepository
}

var _ StorageService = (*Service)(nil)

func (s *Service) SendCoins(sendCoin entities.SendCoin) error {
	fromUser, err := s.AuthRepository.GetUserByUsername(sendCoin.FromUser)
	if err != nil {
		return fmt.Errorf("GetUserByUsername fail: %w", err)
	}

	sendCoin.FromUser = fromUser.ID

	toUser, err := s.AuthRepository.GetUserByUsername(sendCoin.ToUser)
	if err != nil {
		return fmt.Errorf("GetUserByUsername fail: %w", err)
	}

	sendCoin.ToUser = toUser.ID

	err = s.StorageRepository.SendCoins(sendCoin)
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
