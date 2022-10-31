package model

import (
	. "fmt"

	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/system"
)

type (
	RepositoryTransactionHandler func(transaction IRepositoryTransaction) error

	IEntity interface {
		Stringer
		Id() int64
		Payload() string
		SetPayload(string)
		Validate() error
	}

	IRepository interface {
		Name() string
		Migrate() error
		GetSqlDatabase() ISqlDatabase
		SetSqlDatabase(ISqlDatabase)
		Serialize(Pointer, error)
	}

	IRepositoryTransaction interface {
		OnCommit(func())
	}
)
