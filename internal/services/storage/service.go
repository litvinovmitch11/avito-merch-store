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

	GetInventory(userID string) (entities.Inventory, error)
	GetHistory(userID string) (entities.CoinHistory, error)
	GetBalance(userID string) (int, error)
}

type Service struct {
	AuthRepository     authrepo.AuthRepository
	ProductsRepository productsrepo.ProductsRepository
	StorageRepository  storagerepo.StorageRepository
}

var _ StorageService = (*Service)(nil)

func (s *Service) SendCoins(sendCoin entities.SendCoin) error {
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

func (s *Service) BuyMerch(userID, merch string) error {
	product, err := s.ProductsRepository.GetProductByTitle(merch)
	if err != nil {
		return fmt.Errorf("GetProduct fail: %w", err)
	}

	sendCoin := entities.SendCoin{
		FromUser: userID,
		ToUser:   "",
		Amount:   product.Price,
	}

	err = s.StorageRepository.BuyMerch(sendCoin, product)
	if err != nil {
		return fmt.Errorf("SendCoins fail: %w", err)
	}

	return nil
}

func (s *Service) GetInventory(userID string) (entities.Inventory, error) {
	inventory, err := s.StorageRepository.GetInventory(userID)
	if err != nil {
		return entities.Inventory{}, fmt.Errorf("GetBalance fail: %w", err)
	}

	return inventory, nil
}

func (s *Service) GetHistory(userID string) (entities.CoinHistory, error) {
	received, err := s.StorageRepository.GetReceived(userID)
	if err != nil {
		return entities.CoinHistory{}, fmt.Errorf("GetReceived fail: %w", err)
	}

	sent, err := s.StorageRepository.GetSent(userID)
	if err != nil {
		return entities.CoinHistory{}, fmt.Errorf("GetSent fail: %w", err)
	}

	return entities.CoinHistory{
		Received: received,
		Sent:     sent,
	}, nil
}

func (s *Service) GetBalance(userID string) (int, error) {
	balace, err := s.StorageRepository.GetBalance(userID)
	if err != nil {
		return balace.Amount, fmt.Errorf("GetBalance fail: %w", err)
	}

	return balace.Amount, nil
}
