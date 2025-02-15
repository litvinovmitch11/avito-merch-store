package storage

import (
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	storagerepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/storage"
)

type StorageService interface {
	SendCoins(sendCoin entities.SendCoin) error
	BuyMerch(userID, merchID string) error
}

type Service struct {
	StorageRepository storagerepo.StorageRepository
}

var _ StorageService = (*Service)(nil)

func (s *Service) SendCoins(sendCoin entities.SendCoin) error {
	return nil
}

func (s *Service) BuyMerch(userID, merchID string) error {
	return nil
}
