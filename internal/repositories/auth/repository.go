package aauth

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
	ErrUserNotFound = errors.New("user not found")
	ErrPDNotFound   = errors.New("personal data not found")
)

type AuthRepository interface {
	CreateUser(user entities.User, userPD entities.UserPersonalData, balance entities.Balance) error

	GetUserByUsername(username string) (entities.User, error)
	GetPersonalData(userID string) (entities.UserPersonalData, error)
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ AuthRepository = (*Repository)(nil)

func (r *Repository) CreateUser(user entities.User, userPD entities.UserPersonalData, balance entities.Balance) error {
	db, tx, err := r.PostgresqlConnection.CreateTranscation()
	if err != nil {
		return fmt.Errorf("CreateTranscation fail: %w", err)
	}

	u_m := postgresql.UserEntityToUserModel(user)
	u_q := table.Users.
		INSERT(
			table.Users.ID,
			table.Users.Username,
		).
		VALUES(
			u_m.ID,
			u_m.Username,
		)

	err = r.PostgresqlConnection.ExecuteInsertQueryTranscation(db, tx, u_q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQueryTranscation fail: %w", err)
	}

	pd_m := postgresql.UserPDEntityToPDModel(userPD)
	pd_q := table.PersonalData.
		INSERT(
			table.PersonalData.ID,
			table.PersonalData.UserID,
			table.PersonalData.HashedPassword,
		).
		VALUES(
			pd_m.ID,
			pd_m.UserID,
			pd_m.HashedPassword,
		)

	err = r.PostgresqlConnection.ExecuteInsertQueryTranscation(db, tx, pd_q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQueryTranscation fail: %w", err)
	}

	s_m := postgresql.DefaultBalaceEntityToStorageModel(balance)
	s_q := table.Storage.
		INSERT(
			table.Storage.ID,
			table.Storage.UserID,
			table.Storage.Balance,
			table.Storage.MerchInfo,
		).
		VALUES(
			s_m.ID,
			s_m.UserID,
			s_m.Balance,
			s_m.MerchInfo,
		)

	err = r.PostgresqlConnection.ExecuteInsertQueryTranscation(db, tx, s_q)
	if err != nil {
		return fmt.Errorf("ExecuteInsertQueryTranscation fail: %w", err)
	}

	err = r.PostgresqlConnection.FinishTranscation(db, tx)
	if err != nil {
		return fmt.Errorf("FinishTranscation fail: %w", err)
	}

	return nil
}

func (r *Repository) GetUserByUsername(username string) (entities.User, error) {
	q := table.Users.
		SELECT(
			table.Users.ID,
			table.Users.Username,
		).
		WHERE(
			table.Users.Username.EQ(
				postgres.String(username),
			),
		)

	var users []model.Users
	err := r.PostgresqlConnection.ExecuteSelectQuery(q, &users)
	if err != nil {
		return entities.User{}, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	if len(users) == 0 {
		return entities.User{}, ErrUserNotFound
	}

	return postgresql.UserModelToEntity(users[0]), nil
}

func (r *Repository) GetPersonalData(userID string) (entities.UserPersonalData, error) {
	q := table.PersonalData.
		SELECT(
			table.PersonalData.ID,
			table.PersonalData.UserID,
			table.PersonalData.HashedPassword,
		).
		WHERE(
			table.PersonalData.UserID.EQ(
				postgres.String(userID),
			),
		)

	var personalData []model.PersonalData
	err := r.PostgresqlConnection.ExecuteSelectQuery(q, &personalData)
	if err != nil {
		return entities.UserPersonalData{}, fmt.Errorf("ExecuteSelectQuery fail: %w", err)
	}

	if len(personalData) == 0 {
		return entities.UserPersonalData{}, ErrPDNotFound
	}

	return postgresql.PDModelToEntity(personalData[0]), nil
}
