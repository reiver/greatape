package repository

import (
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
)

type systemSchedulesRepository struct {
	baseRepository
}

func newSystemSchedulesRepository(logger ILogger) ISystemSchedulesRepository {
	return &systemSchedulesRepository{
		baseRepository: newBaseRepository("system_schedule", "system_schedules", SystemScheduleEntityType, logger, false),
	}
}

func (repository *systemSchedulesRepository) Add(entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `system_schedules` (`id`, `enabled`, `config`, `editor`) VALUES (?, ?, ?, ?);"
	return repository.database.InsertSingle(query, entity.Id(), entity.Enabled(), entity.Config(), editor)
}

func (repository *systemSchedulesRepository) AddAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "INSERT INTO `system_schedules` (`id`, `enabled`, `config`, `editor`) VALUES (?, ?, ?, ?);"
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Enabled(), entity.Config(), editor)
}

func (repository *systemSchedulesRepository) FetchById(id int64) (ISystemScheduleEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "SELECT `id`, `enabled` = b'1', `config` FROM `system_schedules` WHERE `id` = ? AND `status` = 0;"

	var systemScheduleEntity ISystemScheduleEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id      int64
			enabled bool
			config  string
		)

		if err := cursor.Scan(&id, &enabled, &config); err != nil {
			return err
		}

		systemScheduleEntity = NewSystemScheduleEntity(id, enabled, config)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return systemScheduleEntity, nil
}

func (repository *systemSchedulesRepository) Update(entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `enabled` = ?, `config` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, entity.Enabled(), entity.Config(), editor, entity.Id())
}

func (repository *systemSchedulesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `enabled` = ?, `config` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Enabled(), entity.Config(), editor, entity.Id())
}

func (repository *systemSchedulesRepository) Remove(entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *systemSchedulesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `status` = 1, `editor` = ? WHERE `id` = ?;"
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *systemSchedulesRepository) FetchAll() (SystemScheduleEntities, error) {
	// language=SQL
	query := "SELECT `id`, `enabled` = b'1', `config` FROM `system_schedules` WHERE `id` > 0 AND `status` = 0;"

	var systemScheduleEntities SystemScheduleEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id      int64
			enabled bool
			config  string
		)

		if err := cursor.Scan(&id, &enabled, &config); err != nil {
			return err
		}

		systemScheduleEntities = append(systemScheduleEntities, NewSystemScheduleEntity(id, enabled, config))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return systemScheduleEntities, nil
}

func (repository *systemSchedulesRepository) UpdateEnabled(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `enabled` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *systemSchedulesRepository) UpdateEnabledAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `enabled` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *systemSchedulesRepository) UpdateConfig(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `config` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *systemSchedulesRepository) UpdateConfigAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := "UPDATE `system_schedules` SET `config` = ?, `editor` = ? WHERE `id` = ?;"
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
