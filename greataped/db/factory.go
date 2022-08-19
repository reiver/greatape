package db

import "contracts"

const (
	SqliteStorage contracts.StorageType = 0
)

func CreateStorage(componentType contracts.StorageType) contracts.IStorage {
	switch componentType {
	case SqliteStorage:
		return NewSqliteStorage()
	default:
		panic("unknown_storage_type")
	}
}
