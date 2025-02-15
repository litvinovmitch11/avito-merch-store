package postgresql

import (
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

func UserAuthEntityToUserModel(entity entities.UserAuth) model.Users {
	return model.Users{
		ID:       entity.ID,
		Username: entity.Username,
	}
}

func UserAuthEntityToPDModel(entity entities.UserPersonalData) model.PersonalData {
	return model.PersonalData{
		ID:             entity.ID,
		UserID:         entity.UserID,
		HashedPassword: entity.HashedPassword,
	}
}
