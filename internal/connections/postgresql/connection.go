package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

type PostgresqlConnection interface {
	ExecuteSelectQuery(st postgres.SelectStatement, dest interface{}) error
	ExecuteInsertQuery(st postgres.InsertStatement) error
	ExecuteUpdateQuery(st postgres.UpdateStatement) error

	CreateTranscation() (*sql.DB, *sql.Tx, error)
	FinishTranscation(db *sql.DB, tx *sql.Tx) error

	ExecuteSelectQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.SelectStatement, dest interface{}) error
	ExecuteInsertQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.InsertStatement) error
	ExecuteUpdateQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.UpdateStatement) error
}

type Connection struct{}

var _ PostgresqlConnection = (*Connection)(nil)

func (c *Connection) ExecuteSelectQuery(st postgres.SelectStatement, dest interface{}) error {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("sql.Open fail: %w", err)
	}
	defer db.Close()

	err = st.Query(db, dest)
	if err != nil {
		return fmt.Errorf("st.Query fail: %w", err)
	}

	return nil
}

func (c *Connection) ExecuteInsertQuery(st postgres.InsertStatement) error {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("sql.Open fail: %w", err)
	}
	defer db.Close()

	_, err = st.Exec(db)
	if err != nil {
		return fmt.Errorf("st.Exec fail: %w", err)
	}

	return err
}

func (c *Connection) ExecuteUpdateQuery(st postgres.UpdateStatement) error {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return fmt.Errorf("sql.Open fail: %w", err)
	}
	defer db.Close()

	_, err = st.Exec(db)
	if err != nil {
		return fmt.Errorf("st.Exec fail: %w", err)
	}

	return err
}

func (c *Connection) CreateTranscation() (*sql.DB, *sql.Tx, error) {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "merch_store")

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, nil, fmt.Errorf("sql.Open fail: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		db.Close()

		return nil, nil, fmt.Errorf("db.Begin fail: %w", err)
	}

	return db, tx, nil
}

func (c *Connection) FinishTranscation(db *sql.DB, tx *sql.Tx) error {
	defer db.Close()

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit fail: %w", err)
	}

	return err
}

func (c *Connection) ExecuteSelectQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.SelectStatement, dest interface{}) error {
	err := st.Query(tx, dest)
	if err != nil {
		defer db.Close()

		e := tx.Rollback()
		if e != nil {
			return fmt.Errorf("tx.Rollback fail: %w", err)
		}

		return fmt.Errorf("st.Query fail: %w", err)
	}

	return nil
}

func (c *Connection) ExecuteInsertQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.InsertStatement) error {
	_, err := st.Exec(tx)
	if err != nil {
		defer db.Close()

		e := tx.Rollback()
		if e != nil {
			return fmt.Errorf("tx.Rollback fail: %w", err)
		}

		return fmt.Errorf("st.Exec fail: %w", err)
	}

	return err
}

func (c *Connection) ExecuteUpdateQueryTranscation(db *sql.DB, tx *sql.Tx, st postgres.UpdateStatement) error {
	_, err := st.Exec(tx)
	if err != nil {
		defer db.Close()

		e := tx.Rollback()
		if e != nil {
			return fmt.Errorf("tx.Rollback fail: %w", err)
		}

		return fmt.Errorf("st.Exec fail: %w", err)
	}

	return err
}
