package contracts

import . "github.com/xeronith/diamante/contracts/security"

var SystemSchedulePassThroughFilter = func(ISystemSchedule) bool { return true }

type (
	SystemSchedules               []ISystemSchedule
	SystemScheduleIterator        func(ISystemSchedule)
	SystemScheduleCondition       func(ISystemSchedule) bool
	SystemScheduleFilterPredicate func(ISystemSchedule) bool
	SystemScheduleMapPredicate    func(ISystemSchedule) ISystemSchedule
	SystemScheduleCacheCallback   func()

	ISystemSchedule interface {
		IObject
		// Enabled returns 'Enabled' of this 'SystemSchedule' instance.
		Enabled() bool
		// UpdateEnabled directly updates 'Enabled' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateEnabled(enabled bool, editor Identity)
		// UpdateEnabledAtomic updates 'Enabled' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateEnabledAtomic(transaction ITransaction, enabled bool, editor Identity)
		// Config returns 'Config' of this 'SystemSchedule' instance.
		Config() string
		// UpdateConfig directly updates 'Config' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateConfig(config string, editor Identity)
		// UpdateConfigAtomic updates 'Config' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateConfigAtomic(transaction ITransaction, config string, editor Identity)
	}

	ISystemScheduleCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ISystemSchedule
		Append(systemSchedule ISystemSchedule)
		ForEach(SystemScheduleIterator)
		Reverse() ISystemScheduleCollection
		Array() SystemSchedules
	}

	ISystemScheduleManager interface {
		ISystemComponent
		OnCacheChanged(SystemScheduleCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition SystemScheduleCondition) bool
		ListSystemSchedules(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISystemScheduleCollection
		GetSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		AddSystemSchedule(enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleWithCustomId(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleObject(systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleAtomic(transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleWithCustomIdAtomic(id int64, transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleObjectAtomic(transaction ITransaction, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		Log(enabled bool, config string, source string, editor Identity, payload string)
		UpdateSystemSchedule(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		UpdateSystemScheduleObject(id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		UpdateSystemScheduleAtomic(transaction ITransaction, id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		UpdateSystemScheduleObjectAtomic(transaction ITransaction, id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		AddOrUpdateSystemScheduleObject(id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		AddOrUpdateSystemScheduleObjectAtomic(transaction ITransaction, id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error)
		RemoveSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		RemoveSystemScheduleAtomic(transaction ITransaction, id int64, editor Identity) (ISystemSchedule, error)
		Find(id int64) ISystemSchedule
		ForEach(iterator SystemScheduleIterator)
		Filter(predicate SystemScheduleFilterPredicate) ISystemScheduleCollection
		Map(predicate SystemScheduleMapPredicate) ISystemScheduleCollection
	}
)
