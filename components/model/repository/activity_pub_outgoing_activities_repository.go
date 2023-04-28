package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type activityPubOutgoingActivitiesRepository struct {
	baseRepository
}

func newActivityPubOutgoingActivitiesRepository(logger ILogger) IActivityPubOutgoingActivitiesRepository {
	return &activityPubOutgoingActivitiesRepository{
		baseRepository: newBaseRepository("activity_pub_outgoing_activity", "activity_pub_outgoing_activities", ActivityPubOutgoingActivityEntityType, logger, false),
	}
}

func (repository *activityPubOutgoingActivitiesRepository) Add(entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "activity_pub_outgoing_activities" ("id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	return repository.database.InsertSingle(query, entity.Id(), entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw(), editor)
}

func (repository *activityPubOutgoingActivitiesRepository) AddAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "activity_pub_outgoing_activities" ("id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw(), editor)
}

func (repository *activityPubOutgoingActivitiesRepository) FetchById(id int64) (IActivityPubOutgoingActivityEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw" FROM "activity_pub_outgoing_activities" WHERE "id" = $1 AND "status" = 0;`

	var activityPubOutgoingActivityEntity IActivityPubOutgoingActivityEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id               int64
			identityId       int64
			uniqueIdentifier string
			timestamp        int64
			from             string
			to               string
			content          string
			raw              string
		)

		if err := cursor.Scan(&id, &identityId, &uniqueIdentifier, &timestamp, &from, &to, &content, &raw); err != nil {
			return err
		}

		activityPubOutgoingActivityEntity = NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return activityPubOutgoingActivityEntity, nil
}

func (repository *activityPubOutgoingActivitiesRepository) Update(entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "identity_id" = $1, "unique_identifier" = $2, "timestamp" = $3, "from" = $4, "to" = $5, "content" = $6, "raw" = $7, "editor" = $8 WHERE "id" = $9;`
	return repository.database.UpdateSingle(query, entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw(), editor, entity.Id())
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "identity_id" = $1, "unique_identifier" = $2, "timestamp" = $3, "from" = $4, "to" = $5, "content" = $6, "raw" = $7, "editor" = $8 WHERE "id" = $9;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw(), editor, entity.Id())
}

func (repository *activityPubOutgoingActivitiesRepository) Remove(entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *activityPubOutgoingActivitiesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *activityPubOutgoingActivitiesRepository) FetchAll() (ActivityPubOutgoingActivityEntities, error) {
	// language=SQL
	query := `SELECT "id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw" FROM "activity_pub_outgoing_activities" WHERE "id" > 0 AND "status" = 0;`

	var activityPubOutgoingActivityEntities ActivityPubOutgoingActivityEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id               int64
			identityId       int64
			uniqueIdentifier string
			timestamp        int64
			from             string
			to               string
			content          string
			raw              string
		)

		if err := cursor.Scan(&id, &identityId, &uniqueIdentifier, &timestamp, &from, &to, &content, &raw); err != nil {
			return err
		}

		activityPubOutgoingActivityEntities = append(activityPubOutgoingActivityEntities, NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return activityPubOutgoingActivityEntities, nil
}

func (repository *activityPubOutgoingActivitiesRepository) FetchAllByIdentity(identityId int64) (ActivityPubOutgoingActivityEntities, error) {
	if identityId <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	return repository.FetchAllByDependency("identity_id", identityId)
}

func (repository *activityPubOutgoingActivitiesRepository) FetchAllByDependency(dependencyName string, dependencyId int64) (ActivityPubOutgoingActivityEntities, error) {
	// language=SQL
	query := `SELECT "id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw" FROM "activity_pub_outgoing_activities" WHERE "id" > 0 AND "status" = 0`
	query += ` AND "` + dependencyName + `" = $1;`

	var activityPubOutgoingActivityEntities ActivityPubOutgoingActivityEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id               int64
			identityId       int64
			uniqueIdentifier string
			timestamp        int64
			from             string
			to               string
			content          string
			raw              string
		)

		if err := cursor.Scan(&id, &identityId, &uniqueIdentifier, &timestamp, &from, &to, &content, &raw); err != nil {
			return err
		}

		activityPubOutgoingActivityEntities = append(activityPubOutgoingActivityEntities, NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw))
		return nil
	}, query, dependencyId); err != nil {
		return nil, err
	}

	return activityPubOutgoingActivityEntities, nil
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateUniqueIdentifier(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "unique_identifier" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateUniqueIdentifierAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "unique_identifier" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateTimestamp(id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "timestamp" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateTimestampAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "timestamp" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateFrom(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "from" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateFromAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "from" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateTo(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "to" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateToAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "to" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateContent(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "content" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateContentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "content" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateRaw(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "raw" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubOutgoingActivitiesRepository) UpdateRawAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_outgoing_activities" SET "raw" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
