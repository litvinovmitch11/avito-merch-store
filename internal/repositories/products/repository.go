package products

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/table"
)

type ProductsRepository interface {
	AddProduct(product entities.Product) error
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ ProductsRepository = (*Repository)(nil)

func (r *Repository) AddProduct(product entities.Product) error {
	m := postgresql.ProductEntityToModel(product)
	q := table.Products.
		INSERT(
			table.Products.ID,
			table.Products.Title,
			table.Products.Price,
		).
		VALUES(
			m.ID,
			m.Title,
			m.Price,
		)

	err := r.PostgresqlConnection.ExecuteInsertQuery(q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQuery fail: %w", err)
	}

	return nil
}

// func (r *Repository) GetProduct(product entities.Product) error {
// 	m := postgresql.ProductEntityToModel(product)
// 	q := table.Products.
// 		INSERT(
// 			table.Products.ID,
// 			table.Products.Title,
// 			table.Products.Price,
// 		).
// 		VALUES(
// 			m.ID,
// 			m.Title,
// 			m.Price,
// 		)

// 	err := r.PostgresqlConnection.ExecuteInsertQuery(q)
// 	if err != nil {
// 		return fmt.Errorf("ExecuteInsertQuery fail: %w", err)
// 	}

// 	return nil
// }
