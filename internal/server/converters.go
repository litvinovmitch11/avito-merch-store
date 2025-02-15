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
