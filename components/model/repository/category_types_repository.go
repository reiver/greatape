package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type categoryTypesRepository struct {
	baseRepository
}

func newCategoryTypesRepository(logger ILogger) ICategoryTypesRepository {
	return &categoryTypesRepository{
		baseRepository: newBaseRepository("category_type", "category_types", CategoryTypeEntityType, logger, false),
	}
}

func (repository *categoryTypesRepository) Add(entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `category_types` (`id`, `description`, `editor`) VALUES (?, ?, ?);"
	return repository.database.InsertSingle(query, entity.Id(), entity.Description(), editor)
}

func (repository *categoryTypesRepository) AddAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `category_types` (`id`, `description`, `editor`) VALUES (?, ?, ?);"
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Description(), editor)
}

func (repository *categoryTypesRepository) FetchById(id int64) (ICategoryTypeEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "SELECT `id`, `description` FROM `category_types` WHERE `id` = ? AND `status` = 0;"

	var categoryTypeEntity ICategoryTypeEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id          int64
			description string
		)

		if err := cursor.Scan(&id, &description); err != nil {
			return err
		}

		categoryTypeEntity = NewCategoryTypeEntity(id, description)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return categoryTypeEntity, nil
}

func (repository *categoryTypesRepository) Update(entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `description` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, entity.Description(), editor, entity.Id())
}

func (repository *categoryTypesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `description` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Description(), editor, entity.Id())
}

func (repository *categoryTypesRepository) Remove(entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *categoryTypesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *categoryTypesRepository) FetchAll() (CategoryTypeEntities, error) {
	// language=SQL
	query := "SELECT `id`, `description` FROM `category_types` WHERE `id` > 0 AND `status` = 0;"

	var categoryTypeEntities CategoryTypeEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id          int64
			description string
		)

		if err := cursor.Scan(&id, &description); err != nil {
			return err
		}

		categoryTypeEntities = append(categoryTypeEntities, NewCategoryTypeEntity(id, description))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return categoryTypeEntities, nil
}

func (repository *categoryTypesRepository) UpdateDescription(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `description` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *categoryTypesRepository) UpdateDescriptionAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `category_types` SET `description` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
