package storage

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
	ErrBalanceNotFound = errors.New("balance not found")
	ErrLackOfCoins     = errors.New("lack of coins")
)

type StorageRepository interface {
	SendCoins(sendCoin entities.SendCoin) error
	GetBalance(userID string) (entities.Balance, error)
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ StorageRepository = (*Repository)(nil)

func (r *Repository) SendCoins(sendCoin entities.SendCoin) error {
	// под транзакцией получить баланс

	// под транзакцией удалить деньги у первого
	// под транзакцией удалить дать денег второму
	// записать 1 транзакцию
	// записать 2 транзакцию

	db, tx, err := r.PostgresqlConnection.CreateTranscation()
	if err != nil {
		return fmt.Errorf("CreateTranscation fail: %w", err)
	}

	s_q := table.Storage.
		SELECT(
			table.Storage.ID,
			table.Storage.UserID,
			table.Storage.Balance,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(sendCoin.FromUser),
			),
		)

	var balances []model.Storage
	err = r.PostgresqlConnection.ExecuteSelectQueryTranscation(db, tx, s_q, &balances)
	if err != nil {
		return fmt.Errorf("ExecuteSelectQueryTranscation fail: %w", err)
	}

	if len(balances) == 0 {
		return ErrBalanceNotFound
	}

	balance := postgresql.StorageModelToEntity(balances[0])

	if balance.Amount < sendCoin.Amount {
		return ErrLackOfCoins
	}

	err = r.PostgresqlConnection.FinishTranscation(db, tx)
	if err != nil {
		return fmt.Errorf("FinishTranscation fail: %w", err)
	}

	return nil
}

func (r *Repository) GetBalance(userID string) (entities.Balance, error) {
	q := table.Storage.
		SELECT(
			table.Storage.ID,
			table.Storage.UserID,
			table.Storage.Balance,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(userID),
			),
		)

	var balance []model.Storage
	err := r.PostgresqlConnection.ExecuteSelectQuery(q, &balance)
	if err != nil {
		return entities.Balance{}, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	if len(balance) == 0 {
		return entities.Balance{}, ErrBalanceNotFound
	}

	return postgresql.StorageModelToEntity(balance[0]), nil
}
