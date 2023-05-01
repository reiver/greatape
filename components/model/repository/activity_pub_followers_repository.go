package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type activityPubFollowersRepository struct {
	baseRepository
}

func newActivityPubFollowersRepository(logger ILogger) IActivityPubFollowersRepository {
	return &activityPubFollowersRepository{
		baseRepository: newBaseRepository("activity_pub_follower", "activity_pub_followers", ActivityPubFollowerEntityType, logger, false),
	}
}

func (repository *activityPubFollowersRepository) Add(entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "activity_pub_followers" ("id", "handle", "inbox", "subject", "activity", "accepted", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7);`
	return repository.database.InsertSingle(query, entity.Id(), entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor)
}

func (repository *activityPubFollowersRepository) AddAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "activity_pub_followers" ("id", "handle", "inbox", "subject", "activity", "accepted", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor)
}

func (repository *activityPubFollowersRepository) FetchById(id int64) (IActivityPubFollowerEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "handle", "inbox", "subject", "activity", "accepted" = TRUE FROM "activity_pub_followers" WHERE "id" = $1 AND "status" = 0;`

	var activityPubFollowerEntity IActivityPubFollowerEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id       int64
			handle   string
			inbox    string
			subject  string
			activity string
			accepted bool
		)

		if err := cursor.Scan(&id, &handle, &inbox, &subject, &activity, &accepted); err != nil {
			return err
		}

		activityPubFollowerEntity = NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return activityPubFollowerEntity, nil
}

func (repository *activityPubFollowersRepository) Update(entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "handle" = $1, "inbox" = $2, "subject" = $3, "activity" = $4, "accepted" = $5, "editor" = $6 WHERE "id" = $7;`
	return repository.database.UpdateSingle(query, entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor, entity.Id())
}

func (repository *activityPubFollowersRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "handle" = $1, "inbox" = $2, "subject" = $3, "activity" = $4, "accepted" = $5, "editor" = $6 WHERE "id" = $7;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor, entity.Id())
}

func (repository *activityPubFollowersRepository) Remove(entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *activityPubFollowersRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *activityPubFollowersRepository) FetchAll() (ActivityPubFollowerEntities, error) {
	// language=SQL
	query := `SELECT "id", "handle", "inbox", "subject", "activity", "accepted" = TRUE FROM "activity_pub_followers" WHERE "id" > 0 AND "status" = 0;`

	var activityPubFollowerEntities ActivityPubFollowerEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id       int64
			handle   string
			inbox    string
			subject  string
			activity string
			accepted bool
		)

		if err := cursor.Scan(&id, &handle, &inbox, &subject, &activity, &accepted); err != nil {
			return err
		}

		activityPubFollowerEntities = append(activityPubFollowerEntities, NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return activityPubFollowerEntities, nil
}

func (repository *activityPubFollowersRepository) UpdateHandle(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "handle" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateHandleAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "handle" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateInbox(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "inbox" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateInboxAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "inbox" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateSubject(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "subject" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateSubjectAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "subject" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateActivity(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "activity" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateActivityAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "activity" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateAccepted(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "accepted" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateAcceptedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "activity_pub_followers" SET "accepted" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
