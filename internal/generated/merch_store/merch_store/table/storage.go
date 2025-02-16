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

var Storage = newStorageTable("merch_store", "storage", "")

type storageTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	UserID    postgres.ColumnString
	Balance   postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestamp
	UpdatedAt postgres.ColumnTimestamp
	MerchInfo postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type StorageTable struct {
	storageTable

	EXCLUDED storageTable
}

// AS creates new StorageTable with assigned alias
func (a StorageTable) AS(alias string) *StorageTable {
	return newStorageTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StorageTable with assigned schema name
func (a StorageTable) FromSchema(schemaName string) *StorageTable {
	return newStorageTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StorageTable with assigned table prefix
func (a StorageTable) WithPrefix(prefix string) *StorageTable {
	return newStorageTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StorageTable with assigned table suffix
func (a StorageTable) WithSuffix(suffix string) *StorageTable {
	return newStorageTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStorageTable(schemaName, tableName, alias string) *StorageTable {
	return &StorageTable{
		storageTable: newStorageTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newStorageTableImpl("", "excluded", ""),
	}
}

func newStorageTableImpl(schemaName, tableName, alias string) storageTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		UserIDColumn    = postgres.StringColumn("user_id")
		BalanceColumn   = postgres.IntegerColumn("balance")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		MerchInfoColumn = postgres.StringColumn("merch_info")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, BalanceColumn, CreatedAtColumn, UpdatedAtColumn, MerchInfoColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, BalanceColumn, CreatedAtColumn, UpdatedAtColumn, MerchInfoColumn}
	)

	return storageTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		Balance:   BalanceColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		MerchInfo: MerchInfoColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
