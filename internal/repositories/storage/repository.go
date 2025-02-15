package storage

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/model"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/table"
)

type StorageRepository interface {
	SendCoins(fromID, toID string) error
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ StorageRepository = (*Repository)(nil)

func (r *Repository) SendCoins(fromID, toID string) error {
	// под транзакцией получить баланс
	// под транзакцией удалить деньги у первого
	// под транзакцией удалить дать денег второму

	q := table.Storage.
		SELECT(
			table.Storage.ID,
		)

	var storage model.Storage
	err := r.PostgresqlConnection.ExecuteSelectQuery(q, &storage)
	if err != nil {
		return fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	return nil
}
