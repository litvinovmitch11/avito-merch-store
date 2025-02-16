package storage

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/model"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/table"
)

var (
	ErrStorageDataNotFound = errors.New("storage data not found")
	ErrLackOfCoins         = errors.New("lack of coins")
)

type StorageRepository interface {
	SendCoins(sendCoin entities.SendCoin) error
	BuyMerch(sendCoin entities.SendCoin, product entities.Product) error

	GetBalance(userID string) (entities.Balance, error)
	GetInventory(userID string) (entities.Inventory, error)
	GetReceived(userID string) ([]entities.ReceivedItem, error)
	GetSent(userID string) ([]entities.SentItem, error)
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ StorageRepository = (*Repository)(nil)

func (r *Repository) SendCoins(sendCoin entities.SendCoin) error {
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
		return ErrStorageDataNotFound
	}

	balance := postgresql.StorageModelToEntity(balances[0])

	if balance.Amount < sendCoin.Amount {
		return ErrLackOfCoins
	}

	su_q := table.Storage.
		UPDATE(
			table.Storage.Balance,
		).
		SET(
			balance.Amount - sendCoin.Amount,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(sendCoin.FromUser),
			),
		)

	err = r.PostgresqlConnection.ExecuteUpdateQueryTranscation(db, tx, su_q)
	if err != nil {
		return fmt.Errorf("ExecuteUpdateQueryTranscation fail: %w", err)
	}

	s_q = table.Storage.
		SELECT(
			table.Storage.ID,
			table.Storage.UserID,
			table.Storage.Balance,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(sendCoin.ToUser),
			),
		)

	var balancesTo []model.Storage
	err = r.PostgresqlConnection.ExecuteSelectQueryTranscation(db, tx, s_q, &balancesTo)
	if err != nil {
		return fmt.Errorf("ExecuteSelectQueryTranscation fail: %w", err)
	}

	balanceTo := postgresql.StorageModelToEntity(balancesTo[0])
	fmt.Println(balanceTo)

	su_q = table.Storage.
		UPDATE(
			table.Storage.Balance,
		).
		SET(
			balanceTo.Amount + sendCoin.Amount,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(sendCoin.ToUser),
			),
		)

	err = r.PostgresqlConnection.ExecuteUpdateQueryTranscation(db, tx, su_q)
	if err != nil {
		return fmt.Errorf("ExecuteUpdateQueryTranscation fail: %w", err)
	}

	t_q := table.Transactions.
		INSERT(
			table.Transactions.ID,
			table.Transactions.FromID,
			table.Transactions.ToID,
			table.Transactions.Amount,
		).
		VALUES(
			uuid.NewString(),
			sendCoin.FromUser,
			sendCoin.ToUser,
			sendCoin.Amount,
		)

	err = r.PostgresqlConnection.ExecuteInsertQueryTranscation(db, tx, t_q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQueryTranscation fail: %w", err)
	}

	err = r.PostgresqlConnection.FinishTranscation(db, tx)
	if err != nil {
		return fmt.Errorf("FinishTranscation fail: %w", err)
	}

	return nil
}

func (r *Repository) BuyMerch(sendCoin entities.SendCoin, product entities.Product) error {
	db, tx, err := r.PostgresqlConnection.CreateTranscation()
	if err != nil {
		return fmt.Errorf("CreateTranscation fail: %w", err)
	}

	s_q := table.Storage.
		SELECT(
			table.Storage.ID,
			table.Storage.UserID,
			table.Storage.Balance,
			table.Storage.MerchInfo,
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
		err = r.PostgresqlConnection.FinishTranscation(db, tx)
		if err != nil {
			return fmt.Errorf("FinishTranscation fail: %w", err)
		}

		return ErrStorageDataNotFound
	}

	balance := postgresql.StorageModelToEntity(balances[0])
	inventory, err := postgresql.StorageModelToInventoryMap(balances[0])
	if err != nil {
		return fmt.Errorf("StorageModelToInventoryMap fail: %w", err)
	}

	if balance.Amount < sendCoin.Amount {
		err = r.PostgresqlConnection.FinishTranscation(db, tx)
		if err != nil {
			return fmt.Errorf("FinishTranscation fail: %w", err)
		}

		return ErrLackOfCoins
	}

	inventory[product.Title]++
	storageInventory, err := postgresql.InventoryToStorageModel(inventory)
	if err != nil {
		err = r.PostgresqlConnection.FinishTranscation(db, tx)
		if err != nil {
			return fmt.Errorf("FinishTranscation fail: %w", err)
		}

		return fmt.Errorf("InventoryToStorageModel fail: %w", err)
	}

	su_q := table.Storage.
		UPDATE(
			table.Storage.Balance,
			table.Storage.MerchInfo,
		).
		SET(
			balance.Amount-sendCoin.Amount,
			storageInventory,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(sendCoin.FromUser),
			),
		)

	err = r.PostgresqlConnection.ExecuteUpdateQueryTranscation(db, tx, su_q)
	if err != nil {
		return fmt.Errorf("ExecuteUpdateQueryTranscation fail: %w", err)
	}

	t_q := table.Transactions.
		INSERT(
			table.Transactions.ID,
			table.Transactions.FromID,
			table.Transactions.Amount,
		).
		VALUES(
			uuid.NewString(),
			sendCoin.FromUser,
			sendCoin.Amount,
		)

	err = r.PostgresqlConnection.ExecuteInsertQueryTranscation(db, tx, t_q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQueryTranscation fail: %w", err)
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
		return entities.Balance{}, ErrStorageDataNotFound
	}

	return postgresql.StorageModelToEntity(balance[0]), nil
}

func (r *Repository) GetInventory(userID string) (entities.Inventory, error) {
	s_q := table.Storage.
		SELECT(
			table.Storage.MerchInfo,
		).
		WHERE(
			table.Storage.UserID.EQ(
				postgres.String(userID),
			),
		)

	var storage []model.Storage
	err := r.PostgresqlConnection.ExecuteSelectQuery(s_q, &storage)
	if err != nil {
		return entities.Inventory{}, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	if len(storage) == 0 {
		return entities.Inventory{}, ErrStorageDataNotFound
	}

	inventory, err := postgresql.StorageModelToInventory(storage[0])
	if err != nil {
		return entities.Inventory{}, fmt.Errorf("StorageModelToInventory fail: %w", err)
	}

	return inventory, nil
}

func (r *Repository) GetReceived(userID string) ([]entities.ReceivedItem, error) {
	t_q := table.Transactions.
		SELECT(
			table.Users.Username,
			table.Transactions.Amount,
		).
		FROM(
			table.Transactions.
				LEFT_JOIN(table.Users, table.Users.ID.EQ(table.Transactions.FromID)),
		).
		WHERE(
			table.Transactions.ToID.EQ(
				postgres.String(userID),
			),
		)

	var transactions []struct {
		model.Transactions
		model.Users
	}

	err := r.PostgresqlConnection.ExecuteSelectQuery(t_q, &transactions)
	if err != nil {
		return nil, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	history := postgresql.TransactionsModelToReceived(transactions)

	return history, nil
}

func (r *Repository) GetSent(userID string) ([]entities.SentItem, error) {
	s_q := postgres.
		SELECT(
			table.Users.Username,
			table.Transactions.Amount,
		).
		FROM(
			table.Transactions.
				LEFT_JOIN(table.Users, table.Users.ID.EQ(table.Transactions.ToID)),
		).
		WHERE(
			table.Transactions.FromID.EQ(
				postgres.String(userID),
			),
		)

	var transactions []struct {
		model.Transactions
		model.Users
	}

	err := r.PostgresqlConnection.ExecuteSelectQuery(s_q, &transactions)
	if err != nil {
		return nil, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	history := postgresql.TransactionsModelToSent(transactions)

	return history, nil
}
