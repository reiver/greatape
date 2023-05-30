package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type categoriesRepository struct {
	baseRepository
}

func newCategoriesRepository(logger ILogger) ICategoriesRepository {
	return &categoriesRepository{
		baseRepository: newBaseRepository("category", "categories", CategoryEntityType, logger, false),
	}
}

func (repository *categoriesRepository) Add(entity ICategoryEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "categories" ("id", "category_type_id", "category_id", "title", "description", "editor") VALUES ($1, $2, $3, $4, $5, $6);`
	return repository.database.InsertSingle(query, entity.Id(), entity.CategoryTypeId(), entity.CategoryId(), entity.Title(), entity.Description(), editor)
}

func (repository *categoriesRepository) AddAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "categories" ("id", "category_type_id", "category_id", "title", "description", "editor") VALUES ($1, $2, $3, $4, $5, $6);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.CategoryTypeId(), entity.CategoryId(), entity.Title(), entity.Description(), editor)
}

func (repository *categoriesRepository) FetchById(id int64) (ICategoryEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "category_type_id", "category_id", "title", "description", "sort_order"  FROM "categories" WHERE "id" = $1 AND "status" = 0;`

	var categoryEntity ICategoryEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id             int64
			categoryTypeId int64
			categoryId     int64
			title          string
			description    string
			sortOrder      float32
		)

		if err := cursor.Scan(&id, &categoryTypeId, &categoryId, &title, &description, &sortOrder); err != nil {
			return err
		}

		categoryEntity = NewCategoryEntity(id, categoryTypeId, categoryId, title, description, sortOrder)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return categoryEntity, nil
}

func (repository *categoriesRepository) Update(entity ICategoryEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "category_type_id" = $1, "category_id" = $2, "title" = $3, "description" = $4, "editor" = $5 WHERE "id" = $6;`
	return repository.database.UpdateSingle(query, entity.CategoryTypeId(), entity.CategoryId(), entity.Title(), entity.Description(), editor, entity.Id())
}

func (repository *categoriesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "category_type_id" = $1, "category_id" = $2, "title" = $3, "description" = $4, "editor" = $5 WHERE "id" = $6;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.CategoryTypeId(), entity.CategoryId(), entity.Title(), entity.Description(), editor, entity.Id())
}

func (repository *categoriesRepository) Remove(entity ICategoryEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *categoriesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *categoriesRepository) FetchAll() (CategoryEntities, error) {
	// language=SQL
	query := `SELECT "id", "category_type_id", "category_id", "title", "description", "sort_order"  FROM "categories" WHERE "id" > 0 AND "status" = 0;`

	var categoryEntities CategoryEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id             int64
			categoryTypeId int64
			categoryId     int64
			title          string
			description    string
			sortOrder      float32
		)

		if err := cursor.Scan(&id, &categoryTypeId, &categoryId, &title, &description, &sortOrder); err != nil {
			return err
		}

		categoryEntities = append(categoryEntities, NewCategoryEntity(id, categoryTypeId, categoryId, title, description, sortOrder))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return categoryEntities, nil
}

func (repository *categoriesRepository) FetchAllByCategoryType(categoryTypeId int64) (CategoryEntities, error) {
	if categoryTypeId <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	return repository.FetchAllByDependency("category_type_id", categoryTypeId)
}

func (repository *categoriesRepository) FetchAllByCategory(categoryId int64) (CategoryEntities, error) {
	if categoryId <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	return repository.FetchAllByDependency("category_id", categoryId)
}

func (repository *categoriesRepository) FetchAllByDependency(dependencyName string, dependencyId int64) (CategoryEntities, error) {
	// language=SQL
	query := `SELECT "id", "category_type_id", "category_id", "title", "description", "sort_order"  FROM "categories" WHERE "id" > 0 AND "status" = 0`
	query += ` AND "` + dependencyName + `" = $1;`

	var categoryEntities CategoryEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id             int64
			categoryTypeId int64
			categoryId     int64
			title          string
			description    string
			sortOrder      float32
		)

		if err := cursor.Scan(&id, &categoryTypeId, &categoryId, &title, &description, &sortOrder); err != nil {
			return err
		}

		categoryEntities = append(categoryEntities, NewCategoryEntity(id, categoryTypeId, categoryId, title, description, sortOrder))
		return nil
	}, query, dependencyId); err != nil {
		return nil, err
	}

	return categoryEntities, nil
}

func (repository *categoriesRepository) UpdateTitle(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "title" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *categoriesRepository) UpdateTitleAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "title" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *categoriesRepository) UpdateDescription(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "description" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *categoriesRepository) UpdateDescriptionAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "description" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *categoriesRepository) UpdateSortOrder(id int64, value float32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "sort_order" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *categoriesRepository) UpdateSortOrderAtomic(transaction IRepositoryTransaction, id int64, value float32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "categories" SET "sort_order" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
