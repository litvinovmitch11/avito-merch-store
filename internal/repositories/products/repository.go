package products

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/model"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/table"
)

var (
	ErrProductsNotFound = errors.New("product not found")
)

type ProductsRepository interface {
	AddProduct(product entities.Product) error
	GetProductByTitle(title string) (entities.Product, error)
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

func (r *Repository) GetProductByTitle(title string) (entities.Product, error) {
	q := table.Products.
		SELECT(
			table.Products.ID,
			table.Products.Title,
			table.Products.Price,
		).
		WHERE(
			table.Products.Title.EQ(
				postgres.String(title),
			),
		)

	var products []model.Products
	err := r.PostgresqlConnection.ExecuteSelectQuery(q, &products)
	if err != nil {
		return entities.Product{}, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	if len(products) == 0 {
		return entities.Product{}, ErrProductsNotFound
	}

	return postgresql.ProductModelToEntity(products[0]), nil

}
