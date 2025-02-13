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
