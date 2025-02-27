//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Transactions = newTransactionsTable("merch_store", "transactions", "")

type transactionsTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	FromID    postgres.ColumnString
	ToID      postgres.ColumnString
	Amount    postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestamp
	UpdatedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TransactionsTable struct {
	transactionsTable

	EXCLUDED transactionsTable
}

// AS creates new TransactionsTable with assigned alias
func (a TransactionsTable) AS(alias string) *TransactionsTable {
	return newTransactionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransactionsTable with assigned schema name
func (a TransactionsTable) FromSchema(schemaName string) *TransactionsTable {
	return newTransactionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransactionsTable with assigned table prefix
func (a TransactionsTable) WithPrefix(prefix string) *TransactionsTable {
	return newTransactionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransactionsTable with assigned table suffix
func (a TransactionsTable) WithSuffix(suffix string) *TransactionsTable {
	return newTransactionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransactionsTable(schemaName, tableName, alias string) *TransactionsTable {
	return &TransactionsTable{
		transactionsTable: newTransactionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newTransactionsTableImpl("", "excluded", ""),
	}
}

func newTransactionsTableImpl(schemaName, tableName, alias string) transactionsTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		FromIDColumn    = postgres.StringColumn("from_id")
		ToIDColumn      = postgres.StringColumn("to_id")
		AmountColumn    = postgres.IntegerColumn("amount")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		allColumns      = postgres.ColumnList{IDColumn, FromIDColumn, ToIDColumn, AmountColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns  = postgres.ColumnList{FromIDColumn, ToIDColumn, AmountColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return transactionsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		FromID:    FromIDColumn,
		ToID:      ToIDColumn,
		Amount:    AmountColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
