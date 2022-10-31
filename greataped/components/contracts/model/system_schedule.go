package model

type (
	SystemScheduleEntities []ISystemScheduleEntity

	ISystemScheduleEntity interface {
		IEntity
		Enabled() bool
		Config() string
	}

	ISystemSchedulePipeEntity interface {
		ISystemScheduleEntity
		IPipeEntity
	}

	ISystemSchedulesRepository interface {
		IRepository
		Add(entity ISystemScheduleEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error
		FetchById(editor int64) (ISystemScheduleEntity, error)
		Update(entity ISystemScheduleEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error
		Remove(entity ISystemScheduleEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity ISystemScheduleEntity, editor int64) error
		FetchAll() (SystemScheduleEntities, error)
		UpdateEnabled(id int64, value bool, editor int64) error
		UpdateEnabledAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
		UpdateConfig(id int64, value string, editor int64) error
		UpdateConfigAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
