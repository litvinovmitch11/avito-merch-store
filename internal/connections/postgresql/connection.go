package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

type PostgresqlConnection interface {
	ExecuteInsertQuery(st postgres.InsertStatement) error
	ExecuteSelectQuery(st postgres.SelectStatement, dest interface{}) error
}

type Connection struct {
}

var _ PostgresqlConnection = (*Connection)(nil)

func (c *Connection) ExecuteInsertQuery(st postgres.InsertStatement) error {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("sql.Open fail: %w", err)
	}

	defer db.Close()

	_, err = st.Exec(db)
	if err != nil {
		return fmt.Errorf("sql.Exec fail: %w", err)
	}

	return err
}

func (c *Connection) ExecuteSelectQuery(st postgres.SelectStatement, dest interface{}) error {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("sql.Open fail: %w", err)
	}

	defer db.Close()

	err = st.Query(db, dest)
	if err != nil {
		return fmt.Errorf("sql.Exec fail: %w", err)
	}

	return nil
}
