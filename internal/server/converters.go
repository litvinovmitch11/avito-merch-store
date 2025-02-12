package server

import (
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	api "github.com/litvinovmitch11/avito-merch-store/internal/generated"
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
