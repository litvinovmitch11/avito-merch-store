package server

import (
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/api"
)

func postApiAuthRequestToEntity(request api.AuthRequest) entities.UserAuth {
	return entities.UserAuth{
		Username: request.Username,
		Password: request.Password,
	}
}

func postApiAuthEntityToResponse(entity entities.UserToken) api.AuthResponse {
	return api.AuthResponse{
		Token: entity.Token,
	}
}

func postAdminProductsAddToEntity(request api.ProductAddRequest) entities.Product {
	return entities.Product{
		Title: request.Title,
		Price: request.Price,
	}
}

func postAdminProductsAddEntityToResponse(id string) api.ProductAddResponse {
	return api.ProductAddResponse{
		Id: id,
	}
}

func postSendCoinRequestToEntity(request api.SendCoinRequest) entities.SendCoin {
	return entities.SendCoin{
		ToUser: request.ToUser,
		Amount: request.Amount,
	}
}

func inventoryToServer(items []entities.InventoryItem) []struct {
	Quantity *int    `json:"quantity,omitempty"`
	Type     *string `json:"type,omitempty"`
} {
	server := make([]struct {
		Quantity *int    `json:"quantity,omitempty"`
		Type     *string `json:"type,omitempty"`
	}, 0, len(items))

	for _, item := range items {
		server = append(server, struct {
			Quantity *int    `json:"quantity,omitempty"`
			Type     *string `json:"type,omitempty"`
		}{
			Quantity: &item.Quantity,
			Type:     &item.Type,
		})
	}

	return server
}

func receivedToServer(items []entities.ReceivedItem) []struct {
	Amount   *int    `json:"amount,omitempty"`
	FromUser *string `json:"fromUser,omitempty"`
} {
	server := make([]struct {
		Amount   *int    `json:"amount,omitempty"`
		FromUser *string `json:"fromUser,omitempty"`
	}, 0, len(items))

	for _, item := range items {
		server = append(server, struct {
			Amount   *int    `json:"amount,omitempty"`
			FromUser *string `json:"fromUser,omitempty"`
		}{
			Amount:   &item.Amount,
			FromUser: &item.FromUser,
		})
	}

	return server
}

func sentToServer(items []entities.SentItem) []struct {
	Amount *int    `json:"amount,omitempty"`
	ToUser *string `json:"toUser,omitempty"`
} {
	server := make([]struct {
		Amount *int    `json:"amount,omitempty"`
		ToUser *string `json:"toUser,omitempty"`
	}, 0, len(items))

	for _, item := range items {
		server = append(server, struct {
			Amount *int    `json:"amount,omitempty"`
			ToUser *string `json:"toUser,omitempty"`
		}{
			Amount: &item.Amount,
			ToUser: &item.ToUser,
		})
	}

	return server
}

func historyToServer(history entities.CoinHistory) struct {
	Received *[]struct {
		Amount   *int    `json:"amount,omitempty"`
		FromUser *string `json:"fromUser,omitempty"`
	} `json:"received,omitempty"`
	Sent *[]struct {
		Amount *int    `json:"amount,omitempty"`
		ToUser *string `json:"toUser,omitempty"`
	} `json:"sent,omitempty"`
} {
	received := receivedToServer(history.Received)
	sent := sentToServer(history.Sent)

	return struct {
		Received *[]struct {
			Amount   *int    `json:"amount,omitempty"`
			FromUser *string `json:"fromUser,omitempty"`
		} `json:"received,omitempty"`
		Sent *[]struct {
			Amount *int    `json:"amount,omitempty"`
			ToUser *string `json:"toUser,omitempty"`
		} `json:"sent,omitempty"`
	}{
		Received: &received,
		Sent:     &sent,
	}
}

func getApiInfoEntityToResponse(info entities.Info) api.InfoResponse {
	inventory := inventoryToServer(info.Inventory)
	history := historyToServer(info.CoinHistory)

	return api.InfoResponse{
		Coins:       &info.Coins,
		Inventory:   &inventory,
		CoinHistory: &history,
	}
}
