package db

import (
	. "contracts"
)

const (
	SqliteStorage StorageType = iota
	MySQLStorage
	MariaDBStorage
	PostgreSQLStorage
)

func CreateStorage(componentType StorageType) IStorage {
	switch componentType {
	case SqliteStorage:
		return NewSqliteStorage()
	case MySQLStorage:
		return NewMySQLStorage()
	case MariaDBStorage:
		return NewMySQLStorage()
	case PostgreSQLStorage:
		return NewPostgreSQLStorage()
	default:
		panic("unknown_storage_type")
	}
}
