package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type usersRepository struct {
	baseRepository
}

func newUsersRepository(logger ILogger) IUsersRepository {
	return &usersRepository{
		baseRepository: newBaseRepository("user", "users", UserEntityType, logger, false),
	}
}

func (repository *usersRepository) Add(entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "users" ("id", "github", "editor") VALUES ($1, $2, $3);`
	return repository.database.InsertSingle(query, entity.Id(), entity.Github(), editor)
}

func (repository *usersRepository) AddAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "users" ("id", "github", "editor") VALUES ($1, $2, $3);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Github(), editor)
}

func (repository *usersRepository) FetchById(id int64) (IUserEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "github" FROM "users" WHERE "id" = $1 AND "status" = 0;`

	var userEntity IUserEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id     int64
			github string
		)

		if err := cursor.Scan(&id, &github); err != nil {
			return err
		}

		userEntity = NewUserEntity(id, github)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (repository *usersRepository) Update(entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "github" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, entity.Github(), editor, entity.Id())
}

func (repository *usersRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "github" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Github(), editor, entity.Id())
}

func (repository *usersRepository) Remove(entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *usersRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *usersRepository) FetchAll() (UserEntities, error) {
	// language=SQL
	query := `SELECT "id", "github" FROM "users" WHERE "id" > 0 AND "status" = 0;`

	var userEntities UserEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id     int64
			github string
		)

		if err := cursor.Scan(&id, &github); err != nil {
			return err
		}

		userEntities = append(userEntities, NewUserEntity(id, github))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return userEntities, nil
}

func (repository *usersRepository) UpdateGithub(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "github" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *usersRepository) UpdateGithubAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "users" SET "github" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
