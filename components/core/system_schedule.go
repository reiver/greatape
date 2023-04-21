package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type systemSchedule struct {
	object
	enabled bool
	config  string
}

// noinspection GoUnusedExportedFunction
func InitializeSystemSchedule() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error) {
	instance := &systemSchedule{
		object: object{
			id: id,
		},
		enabled: enabled,
		config:  config,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewSystemScheduleFromEntity(entity ISystemScheduleEntity) (ISystemSchedule, error) {
	instance := &systemSchedule{
		object: object{
			id: entity.Id(),
		},
		enabled: entity.Enabled(),
		config:  entity.Config(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (systemSchedule *systemSchedule) Enabled() bool {
	return systemSchedule.enabled
}

func (systemSchedule *systemSchedule) UpdateEnabled(enabled bool, editor Identity) {
	if err := repository.SystemSchedules.UpdateEnabled(systemSchedule.id, enabled, editor.Id()); err != nil {
		panic(err.Error())
	}

	systemSchedule.enabled = enabled
}

func (systemSchedule *systemSchedule) UpdateEnabledAtomic(transaction ITransaction, enabled bool, editor Identity) {
	transaction.OnCommit(func() {
		systemSchedule.enabled = enabled
	})

	if err := repository.SystemSchedules.UpdateEnabledAtomic(transaction, systemSchedule.id, enabled, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (systemSchedule *systemSchedule) Config() string {
	return systemSchedule.config
}

func (systemSchedule *systemSchedule) UpdateConfig(config string, editor Identity) {
	if err := repository.SystemSchedules.UpdateConfig(systemSchedule.id, config, editor.Id()); err != nil {
		panic(err.Error())
	}

	systemSchedule.config = config
}

func (systemSchedule *systemSchedule) UpdateConfigAtomic(transaction ITransaction, config string, editor Identity) {
	transaction.OnCommit(func() {
		systemSchedule.config = config
	})

	if err := repository.SystemSchedules.UpdateConfigAtomic(transaction, systemSchedule.id, config, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (systemSchedule *systemSchedule) Validate() error {
	return nil
}

func (systemSchedule *systemSchedule) String() string {
	return fmt.Sprintf("SystemSchedule (Id: %d, Enabled: %v, Config: %v)", systemSchedule.Id(), systemSchedule.Enabled(), systemSchedule.Config())
}

//------------------------------------------------------------------------------

type systemSchedules struct {
	collection SystemSchedules
}

// NewSystemSchedules creates an empty collection of 'System Schedule' which is not thread-safe.
func NewSystemSchedules() ISystemScheduleCollection {
	return &systemSchedules{
		collection: make(SystemSchedules, 0),
	}
}

func (systemSchedules *systemSchedules) Count() int {
	return len(systemSchedules.collection)
}

func (systemSchedules *systemSchedules) IsEmpty() bool {
	return len(systemSchedules.collection) == 0
}

func (systemSchedules *systemSchedules) IsNotEmpty() bool {
	return len(systemSchedules.collection) > 0
}

func (systemSchedules *systemSchedules) HasExactlyOneItem() bool {
	return len(systemSchedules.collection) == 1
}

func (systemSchedules *systemSchedules) HasAtLeastOneItem() bool {
	return len(systemSchedules.collection) >= 1
}

func (systemSchedules *systemSchedules) First() ISystemSchedule {
	return systemSchedules.collection[0]
}

func (systemSchedules *systemSchedules) Append(systemSchedule ISystemSchedule) {
	systemSchedules.collection = append(systemSchedules.collection, systemSchedule)
}

func (systemSchedules *systemSchedules) ForEach(iterator SystemScheduleIterator) {
	if iterator == nil {
		return
	}

	for _, value := range systemSchedules.collection {
		iterator(value)
	}
}

func (systemSchedules *systemSchedules) Array() SystemSchedules {
	return systemSchedules.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) SystemScheduleExists(id int64) bool {
	return dispatcher.conductor.SystemScheduleManager().Exists(id)
}

func (dispatcher *dispatcher) SystemScheduleExistsWhich(condition SystemScheduleCondition) bool {
	return dispatcher.conductor.SystemScheduleManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListSystemSchedules() ISystemScheduleCollection {
	return dispatcher.conductor.SystemScheduleManager().ListSystemSchedules(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachSystemSchedule(iterator SystemScheduleIterator) {
	dispatcher.conductor.SystemScheduleManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterSystemSchedules(predicate SystemScheduleFilterPredicate) ISystemScheduleCollection {
	return dispatcher.conductor.SystemScheduleManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapSystemSchedules(predicate SystemScheduleMapPredicate) ISystemScheduleCollection {
	return dispatcher.conductor.SystemScheduleManager().Map(predicate)
}

func (dispatcher *dispatcher) GetSystemSchedule(id int64) ISystemSchedule {
	if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().GetSystemSchedule(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return systemSchedule
	}
}

func (dispatcher *dispatcher) AddSystemSchedule(enabled bool, config string) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddSystemScheduleAtomic(transaction, enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddSystemSchedule(enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}

func (dispatcher *dispatcher) AddSystemScheduleWithCustomId(id int64, enabled bool, config string) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddSystemScheduleWithCustomIdAtomic(id, transaction, enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddSystemScheduleWithCustomId(id, enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}

func (dispatcher *dispatcher) LogSystemSchedule(enabled bool, config string, source string, payload string) {
	dispatcher.conductor.SystemScheduleManager().Log(enabled, config, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateSystemSchedule(id int64, enabled bool, config string) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().UpdateSystemScheduleAtomic(transaction, id, enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().UpdateSystemSchedule(id, enabled, config, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateSystemScheduleObject(object IObject, systemSchedule ISystemSchedule) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().UpdateSystemScheduleAtomic(transaction, object.Id(), systemSchedule.Enabled(), systemSchedule.Config(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().UpdateSystemSchedule(object.Id(), systemSchedule.Enabled(), systemSchedule.Config(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateSystemScheduleObject(object IObject, systemSchedule ISystemSchedule) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddOrUpdateSystemScheduleObjectAtomic(transaction, object.Id(), systemSchedule, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().AddOrUpdateSystemScheduleObject(object.Id(), systemSchedule, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}

func (dispatcher *dispatcher) RemoveSystemSchedule(id int64) ISystemSchedule {
	transaction := dispatcher.transaction
	if transaction != nil {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().RemoveSystemScheduleAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	} else {
		if systemSchedule, err := dispatcher.conductor.SystemScheduleManager().RemoveSystemSchedule(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return systemSchedule
		}
	}
}
