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

type accessControl struct {
	object
	key   uint64
	value uint64
}

// noinspection GoUnusedExportedFunction
func InitializeAccessControl() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error) {
	instance := &accessControl{
		object: object{
			id: id,
		},
		key:   key,
		value: value,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewAccessControlFromEntity(entity IAccessControlEntity) (IAccessControl, error) {
	instance := &accessControl{
		object: object{
			id: entity.Id(),
		},
		key:   entity.Key(),
		value: entity.Value(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (accessControl *accessControl) Key() uint64 {
	return accessControl.key
}

func (accessControl *accessControl) UpdateKey(key uint64, editor Identity) {
	if err := repository.AccessControls.UpdateKey(accessControl.id, key, editor.Id()); err != nil {
		panic(err.Error())
	}

	accessControl.key = key
}

func (accessControl *accessControl) UpdateKeyAtomic(transaction ITransaction, key uint64, editor Identity) {
	transaction.OnCommit(func() {
		accessControl.key = key
	})

	if err := repository.AccessControls.UpdateKeyAtomic(transaction, accessControl.id, key, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (accessControl *accessControl) Value() uint64 {
	return accessControl.value
}

func (accessControl *accessControl) UpdateValue(value uint64, editor Identity) {
	if err := repository.AccessControls.UpdateValue(accessControl.id, value, editor.Id()); err != nil {
		panic(err.Error())
	}

	accessControl.value = value
}

func (accessControl *accessControl) UpdateValueAtomic(transaction ITransaction, value uint64, editor Identity) {
	transaction.OnCommit(func() {
		accessControl.value = value
	})

	if err := repository.AccessControls.UpdateValueAtomic(transaction, accessControl.id, value, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (accessControl *accessControl) Validate() error {
	return nil
}

func (accessControl *accessControl) String() string {
	return fmt.Sprintf("AccessControl (Id: %d, Key: %v, Value: %v)", accessControl.Id(), accessControl.Key(), accessControl.Value())
}

//------------------------------------------------------------------------------

type accessControls struct {
	collection AccessControls
}

// NewAccessControls creates an empty collection of 'Access Control' which is not thread-safe.
func NewAccessControls() IAccessControlCollection {
	return &accessControls{
		collection: make(AccessControls, 0),
	}
}

func (accessControls *accessControls) Count() int {
	return len(accessControls.collection)
}

func (accessControls *accessControls) IsEmpty() bool {
	return len(accessControls.collection) == 0
}

func (accessControls *accessControls) IsNotEmpty() bool {
	return len(accessControls.collection) > 0
}

func (accessControls *accessControls) HasExactlyOneItem() bool {
	return len(accessControls.collection) == 1
}

func (accessControls *accessControls) HasAtLeastOneItem() bool {
	return len(accessControls.collection) >= 1
}

func (accessControls *accessControls) First() IAccessControl {
	return accessControls.collection[0]
}

func (accessControls *accessControls) Append(accessControl IAccessControl) {
	accessControls.collection = append(accessControls.collection, accessControl)
}

func (accessControls *accessControls) Reverse() IAccessControlCollection {
	slice := accessControls.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	accessControls.collection = slice

	return accessControls
}

func (accessControls *accessControls) ForEach(iterator AccessControlIterator) {
	if iterator == nil {
		return
	}

	for _, value := range accessControls.collection {
		iterator(value)
	}
}

func (accessControls *accessControls) Array() AccessControls {
	return accessControls.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) AccessControlExists(id int64) bool {
	return dispatcher.conductor.AccessControlManager().Exists(id)
}

func (dispatcher *dispatcher) AccessControlExistsWhich(condition AccessControlCondition) bool {
	return dispatcher.conductor.AccessControlManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListAccessControls() IAccessControlCollection {
	return dispatcher.conductor.AccessControlManager().ListAccessControls(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachAccessControl(iterator AccessControlIterator) {
	dispatcher.conductor.AccessControlManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterAccessControls(predicate AccessControlFilterPredicate) IAccessControlCollection {
	return dispatcher.conductor.AccessControlManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapAccessControls(predicate AccessControlMapPredicate) IAccessControlCollection {
	return dispatcher.conductor.AccessControlManager().Map(predicate)
}

func (dispatcher *dispatcher) GetAccessControl(id int64) IAccessControl {
	if accessControl, err := dispatcher.conductor.AccessControlManager().GetAccessControl(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return accessControl
	}
}

func (dispatcher *dispatcher) AddAccessControl(key uint64, value uint64) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddAccessControlAtomic(transaction, key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddAccessControl(key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}

func (dispatcher *dispatcher) AddAccessControlWithCustomId(id int64, key uint64, value uint64) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddAccessControlWithCustomIdAtomic(id, transaction, key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddAccessControlWithCustomId(id, key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}

func (dispatcher *dispatcher) LogAccessControl(key uint64, value uint64, source string, payload string) {
	dispatcher.conductor.AccessControlManager().Log(key, value, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateAccessControl(id int64, key uint64, value uint64) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().UpdateAccessControlAtomic(transaction, id, key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().UpdateAccessControl(id, key, value, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateAccessControlObject(object IObject, accessControl IAccessControl) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().UpdateAccessControlAtomic(transaction, object.Id(), accessControl.Key(), accessControl.Value(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().UpdateAccessControl(object.Id(), accessControl.Key(), accessControl.Value(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateAccessControlObject(object IObject, accessControl IAccessControl) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddOrUpdateAccessControlObjectAtomic(transaction, object.Id(), accessControl, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().AddOrUpdateAccessControlObject(object.Id(), accessControl, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}

func (dispatcher *dispatcher) RemoveAccessControl(id int64) IAccessControl {
	transaction := dispatcher.transaction
	if transaction != nil {
		if accessControl, err := dispatcher.conductor.AccessControlManager().RemoveAccessControlAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	} else {
		if accessControl, err := dispatcher.conductor.AccessControlManager().RemoveAccessControl(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return accessControl
		}
	}
}
