package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type accessControlsRepository struct {
	baseRepository
}

func newAccessControlsRepository(logger ILogger) IAccessControlsRepository {
	return &accessControlsRepository{
		baseRepository: newBaseRepository("access_control", "access_controls", AccessControlEntityType, logger, false),
	}
}

func (repository *accessControlsRepository) Add(entity IAccessControlEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "access_controls" ("id", "key", "value", "editor") VALUES ($1, $2, $3, $4);`
	return repository.database.InsertSingle(query, entity.Id(), entity.Key(), entity.Value(), editor)
}

func (repository *accessControlsRepository) AddAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "access_controls" ("id", "key", "value", "editor") VALUES ($1, $2, $3, $4);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Key(), entity.Value(), editor)
}

func (repository *accessControlsRepository) FetchById(id int64) (IAccessControlEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "key", "value" FROM "access_controls" WHERE "id" = $1 AND "status" = 0;`

	var accessControlEntity IAccessControlEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id    int64
			key   uint64
			value uint64
		)

		if err := cursor.Scan(&id, &key, &value); err != nil {
			return err
		}

		accessControlEntity = NewAccessControlEntity(id, key, value)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return accessControlEntity, nil
}

func (repository *accessControlsRepository) Update(entity IAccessControlEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "key" = $1, "value" = $2, "editor" = $3 WHERE "id" = $4;`
	return repository.database.UpdateSingle(query, entity.Key(), entity.Value(), editor, entity.Id())
}

func (repository *accessControlsRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "key" = $1, "value" = $2, "editor" = $3 WHERE "id" = $4;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Key(), entity.Value(), editor, entity.Id())
}

func (repository *accessControlsRepository) Remove(entity IAccessControlEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *accessControlsRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *accessControlsRepository) FetchAll() (AccessControlEntities, error) {
	// language=SQL
	query := `SELECT "id", "key", "value" FROM "access_controls" WHERE "id" > 0 AND "status" = 0;`

	var accessControlEntities AccessControlEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id    int64
			key   uint64
			value uint64
		)

		if err := cursor.Scan(&id, &key, &value); err != nil {
			return err
		}

		accessControlEntities = append(accessControlEntities, NewAccessControlEntity(id, key, value))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return accessControlEntities, nil
}

func (repository *accessControlsRepository) UpdateKey(id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *accessControlsRepository) UpdateKeyAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *accessControlsRepository) UpdateValue(id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "value" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *accessControlsRepository) UpdateValueAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "access_controls" SET "value" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
