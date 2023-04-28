package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type remoteActivitiesRepository struct {
	baseRepository
}

func newRemoteActivitiesRepository(logger ILogger) IRemoteActivitiesRepository {
	return &remoteActivitiesRepository{
		baseRepository: newBaseRepository("remote_activity", "remote_activities", RemoteActivityEntityType, logger, false),
	}
}

func (repository *remoteActivitiesRepository) Add(entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "remote_activities" ("id", "entry_point", "duration", "successful", "error_message", "remote_address", "user_agent", "event_type", "timestamp", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	return repository.database.InsertSingle(query, entity.Id(), entity.EntryPoint(), entity.Duration(), entity.Successful(), entity.ErrorMessage(), entity.RemoteAddress(), entity.UserAgent(), entity.EventType(), entity.Timestamp(), editor)
}

func (repository *remoteActivitiesRepository) AddAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "remote_activities" ("id", "entry_point", "duration", "successful", "error_message", "remote_address", "user_agent", "event_type", "timestamp", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.EntryPoint(), entity.Duration(), entity.Successful(), entity.ErrorMessage(), entity.RemoteAddress(), entity.UserAgent(), entity.EventType(), entity.Timestamp(), editor)
}

func (repository *remoteActivitiesRepository) FetchById(id int64) (IRemoteActivityEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "entry_point", "duration", "successful" = TRUE, "error_message", "remote_address", "user_agent", "event_type", "timestamp" FROM "remote_activities" WHERE "id" = $1 AND "status" = 0;`

	var remoteActivityEntity IRemoteActivityEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id            int64
			entryPoint    string
			duration      int64
			successful    bool
			errorMessage  string
			remoteAddress string
			userAgent     string
			eventType     uint32
			timestamp     int64
		)

		if err := cursor.Scan(&id, &entryPoint, &duration, &successful, &errorMessage, &remoteAddress, &userAgent, &eventType, &timestamp); err != nil {
			return err
		}

		remoteActivityEntity = NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return remoteActivityEntity, nil
}

func (repository *remoteActivitiesRepository) Update(entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "entry_point" = $1, "duration" = $2, "successful" = $3, "error_message" = $4, "remote_address" = $5, "user_agent" = $6, "event_type" = $7, "timestamp" = $8, "editor" = $9 WHERE "id" = $10;`
	return repository.database.UpdateSingle(query, entity.EntryPoint(), entity.Duration(), entity.Successful(), entity.ErrorMessage(), entity.RemoteAddress(), entity.UserAgent(), entity.EventType(), entity.Timestamp(), editor, entity.Id())
}

func (repository *remoteActivitiesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "entry_point" = $1, "duration" = $2, "successful" = $3, "error_message" = $4, "remote_address" = $5, "user_agent" = $6, "event_type" = $7, "timestamp" = $8, "editor" = $9 WHERE "id" = $10;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.EntryPoint(), entity.Duration(), entity.Successful(), entity.ErrorMessage(), entity.RemoteAddress(), entity.UserAgent(), entity.EventType(), entity.Timestamp(), editor, entity.Id())
}

func (repository *remoteActivitiesRepository) Remove(entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *remoteActivitiesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *remoteActivitiesRepository) FetchAll() (RemoteActivityEntities, error) {
	// language=SQL
	query := `SELECT "id", "entry_point", "duration", "successful" = TRUE, "error_message", "remote_address", "user_agent", "event_type", "timestamp" FROM "remote_activities" WHERE "id" > 0 AND "status" = 0;`

	var remoteActivityEntities RemoteActivityEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id            int64
			entryPoint    string
			duration      int64
			successful    bool
			errorMessage  string
			remoteAddress string
			userAgent     string
			eventType     uint32
			timestamp     int64
		)

		if err := cursor.Scan(&id, &entryPoint, &duration, &successful, &errorMessage, &remoteAddress, &userAgent, &eventType, &timestamp); err != nil {
			return err
		}

		remoteActivityEntities = append(remoteActivityEntities, NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return remoteActivityEntities, nil
}

func (repository *remoteActivitiesRepository) UpdateEntryPoint(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "entry_point" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateEntryPointAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "entry_point" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateDuration(id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "duration" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateDurationAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "duration" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateSuccessful(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "successful" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateSuccessfulAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "successful" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateErrorMessage(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "error_message" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateErrorMessageAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "error_message" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateRemoteAddress(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "remote_address" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateRemoteAddressAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "remote_address" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateUserAgent(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "user_agent" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateUserAgentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "user_agent" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateEventType(id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "event_type" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateEventTypeAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "event_type" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateTimestamp(id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "timestamp" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *remoteActivitiesRepository) UpdateTimestampAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "remote_activities" SET "timestamp" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
