package repository

import (
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
)

type documentsRepository struct {
	baseRepository
}

func newDocumentsRepository(logger ILogger) IDocumentsRepository {
	return &documentsRepository{
		baseRepository: newBaseRepository("document", "documents", DocumentEntityType, logger, false),
	}
}

func (repository *documentsRepository) Add(entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `documents` (`id`, `content`, `editor`) VALUES (?, ?, ?);"
	return repository.database.InsertSingle(query, entity.Id(), entity.Content(), editor)
}

func (repository *documentsRepository) AddAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `documents` (`id`, `content`, `editor`) VALUES (?, ?, ?);"
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Content(), editor)
}

func (repository *documentsRepository) FetchById(id int64) (IDocumentEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "SELECT `id`, `content` FROM `documents` WHERE `id` = ? AND `status` = 0;"

	var documentEntity IDocumentEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id      int64
			content string
		)

		if err := cursor.Scan(&id, &content); err != nil {
			return err
		}

		documentEntity = NewDocumentEntity(id, content)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return documentEntity, nil
}

func (repository *documentsRepository) Update(entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `content` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, entity.Content(), editor, entity.Id())
}

func (repository *documentsRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `content` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Content(), editor, entity.Id())
}

func (repository *documentsRepository) Remove(entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *documentsRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *documentsRepository) FetchAll() (DocumentEntities, error) {
	// language=SQL
	query := "SELECT `id`, `content` FROM `documents` WHERE `id` > 0 AND `status` = 0;"

	var documentEntities DocumentEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id      int64
			content string
		)

		if err := cursor.Scan(&id, &content); err != nil {
			return err
		}

		documentEntities = append(documentEntities, NewDocumentEntity(id, content))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return documentEntities, nil
}

func (repository *documentsRepository) UpdateContent(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `content` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *documentsRepository) UpdateContentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `documents` SET `content` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
