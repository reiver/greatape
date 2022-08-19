package db

import (
	. "contracts"

	"gorm.io/gorm"
)

var Executor *gorm.DB

const (
	SqliteStorage StorageType = 0
	MySQLStorage  StorageType = 1
)

func CreateStorage(componentType StorageType) IStorage {
	switch componentType {
	case SqliteStorage:
		return NewSqliteStorage()
	case MySQLStorage:
		return NewMySQLStorage()
	default:
		panic("unknown_storage_type")
	}
}
