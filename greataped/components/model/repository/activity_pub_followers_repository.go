package repository

import (
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
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
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `activity_pub_followers` (`id`, `handle`, `inbox`, `subject`, `activity`, `accepted`, `editor`) VALUES (?, ?, ?, ?, ?, ?, ?);"
	return repository.database.InsertSingle(query, entity.Id(), entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor)
}

func (repository *activityPubFollowersRepository) AddAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `activity_pub_followers` (`id`, `handle`, `inbox`, `subject`, `activity`, `accepted`, `editor`) VALUES (?, ?, ?, ?, ?, ?, ?);"
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor)
}

func (repository *activityPubFollowersRepository) FetchById(id int64) (IActivityPubFollowerEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "SELECT `id`, `handle`, `inbox`, `subject`, `activity`, `accepted` = b'1' FROM `activity_pub_followers` WHERE `id` = ? AND `status` = 0;"

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
	query := "UPDATE `activity_pub_followers` SET `handle` = ?, `inbox` = ?, `subject` = ?, `activity` = ?, `accepted` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor, entity.Id())
}

func (repository *activityPubFollowersRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `handle` = ?, `inbox` = ?, `subject` = ?, `activity` = ?, `accepted` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted(), editor, entity.Id())
}

func (repository *activityPubFollowersRepository) Remove(entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *activityPubFollowersRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *activityPubFollowersRepository) FetchAll() (ActivityPubFollowerEntities, error) {
	// language=SQL
	query := "SELECT `id`, `handle`, `inbox`, `subject`, `activity`, `accepted` = b'1' FROM `activity_pub_followers` WHERE `id` > 0 AND `status` = 0;"

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
	query := "UPDATE `activity_pub_followers` SET `handle` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateHandleAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `handle` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateInbox(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `inbox` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateInboxAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `inbox` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateSubject(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `subject` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateSubjectAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `subject` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateActivity(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `activity` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateActivityAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `activity` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateAccepted(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `accepted` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *activityPubFollowersRepository) UpdateAcceptedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `activity_pub_followers` SET `accepted` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
