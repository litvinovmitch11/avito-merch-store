package postgresql

import (
	"encoding/json"
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/model"
)

func ProductEntityToModel(entity entities.Product) model.Products {
	return model.Products{
		ID:    entity.Id,
		Title: entity.Title,
		Price: int32(entity.Price),
	}
}

func ProductModelToEntity(model model.Products) entities.Product {
	return entities.Product{
		Id:    model.ID,
		Title: model.Title,
		Price: int(model.Price),
	}
}

func UserEntityToUserModel(entity entities.User) model.Users {
	return model.Users{
		ID:       entity.ID,
		Username: entity.Username,
	}
}

func UserModelToEntity(model model.Users) entities.User {
	return entities.User{
		ID:       model.ID,
		Username: model.Username,
	}
}

func UserPDEntityToPDModel(entity entities.UserPersonalData) model.PersonalData {
	return model.PersonalData{
		ID:             entity.ID,
		UserID:         entity.UserID,
		HashedPassword: entity.HashedPassword,
	}
}

func PDModelToEntity(model model.PersonalData) entities.UserPersonalData {
	return entities.UserPersonalData{
		ID:             model.ID,
		UserID:         model.UserID,
		HashedPassword: model.HashedPassword,
	}
}

func DefaultBalaceEntityToStorageModel(balance entities.Balance) model.Storage {
	return model.Storage{
		ID:        balance.ID,
		UserID:    balance.UserID,
		Balance:   int32(balance.Amount),
		MerchInfo: "[]",
	}
}

func StorageModelToEntity(model model.Storage) entities.Balance {
	return entities.Balance{
		ID:     model.ID,
		UserID: model.UserID,
		Amount: int(model.Balance),
	}
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func StorageModelToInventory(model model.Storage) (map[string]int, error) {
	var inventory []InventoryItem

	err := json.Unmarshal([]byte(model.MerchInfo), &inventory)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal fail: %w", err)
	}

	inventoryMap := make(map[string]int)
	for _, item := range inventory {
		inventoryMap[item.Type] = item.Quantity
	}

	return inventoryMap, nil
}

func InventoryToStorageModel(inventoryMap map[string]int) (string, error) {
	inventory := make([]InventoryItem, 0, len(inventoryMap))

	for t, q := range inventoryMap {
		inventory = append(inventory, InventoryItem{Type: t, Quantity: q})
	}

	bytes, err := json.Marshal(inventory)
	if err != nil {
		return "", fmt.Errorf("json.Marshal fail: %w", err)
	}

	return string(bytes), nil
}
