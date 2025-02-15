package aauth

import (
	"fmt"

	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/merch_store/merch_store/table"
)

type AuthRepository interface {
	CreateUser(userAuth entities.UserAuth, userPD entities.UserPersonalData) error
}

type Repository struct {
	PostgresqlConnection postgresql.PostgresqlConnection
}

var _ AuthRepository = (*Repository)(nil)

func (r *Repository) CreateUser(userAuth entities.UserAuth, userPD entities.UserPersonalData) error {
	db, tx, err := r.PostgresqlConnection.CreateTranscation()
	if err != nil {
		return fmt.Errorf("CreateTranscation fail: %w", err)
	}

	u_m := postgresql.UserAuthEntityToUserModel(userAuth)
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
		return fmt.Errorf("ExecuteInsertQuery fail: %w", err)
	}

	pd_m := postgresql.UserAuthEntityToPDModel(userPD)
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
		return fmt.Errorf("ExecuteInsertQuery fail: %w", err)
	}

	err = r.PostgresqlConnection.FinishTranscation(db, tx)
	if err != nil {
		return fmt.Errorf("FinishTranscation fail: %w", err)
	}

	return nil
}
