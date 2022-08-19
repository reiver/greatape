package db

import (
	. "contracts"

	"gorm.io/gorm"
)

var Executor *gorm.DB

const (
	SqliteStorage     StorageType = 0
	MySQLStorage      StorageType = 1
	MariaDBStorage    StorageType = 2
	PostgreSQLStorage StorageType = 3
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
