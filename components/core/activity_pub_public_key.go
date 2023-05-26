package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type activityPubPublicKey struct {
	id           string
	owner        string
	publicKeyPem string
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubPublicKey() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubPublicKey() (IActivityPubPublicKey, error) {
	instance := &activityPubPublicKey{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubPublicKey *activityPubPublicKey) Id() string {
	return activityPubPublicKey.id
}

func (activityPubPublicKey *activityPubPublicKey) SetId(value string) {
	activityPubPublicKey.id = value
}

func (activityPubPublicKey *activityPubPublicKey) Owner() string {
	return activityPubPublicKey.owner
}

func (activityPubPublicKey *activityPubPublicKey) SetOwner(value string) {
	activityPubPublicKey.owner = value
}

func (activityPubPublicKey *activityPubPublicKey) PublicKeyPem() string {
	return activityPubPublicKey.publicKeyPem
}

func (activityPubPublicKey *activityPubPublicKey) SetPublicKeyPem(value string) {
	activityPubPublicKey.publicKeyPem = value
}

func (activityPubPublicKey *activityPubPublicKey) Validate() error {
	return nil
}

func (activityPubPublicKey *activityPubPublicKey) String() string {
	return fmt.Sprintf("ActivityPubPublicKey (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type activityPubPublicKeys struct {
	collection ActivityPubPublicKeys
}

// NewActivityPubPublicKeys creates an empty collection of 'Activity Pub Public Key' which is not thread-safe.
func NewActivityPubPublicKeys() IActivityPubPublicKeyCollection {
	return &activityPubPublicKeys{
		collection: make(ActivityPubPublicKeys, 0),
	}
}

func (activityPubPublicKeys *activityPubPublicKeys) Count() int {
	return len(activityPubPublicKeys.collection)
}

func (activityPubPublicKeys *activityPubPublicKeys) IsEmpty() bool {
	return len(activityPubPublicKeys.collection) == 0
}

func (activityPubPublicKeys *activityPubPublicKeys) IsNotEmpty() bool {
	return len(activityPubPublicKeys.collection) > 0
}

func (activityPubPublicKeys *activityPubPublicKeys) HasExactlyOneItem() bool {
	return len(activityPubPublicKeys.collection) == 1
}

func (activityPubPublicKeys *activityPubPublicKeys) HasAtLeastOneItem() bool {
	return len(activityPubPublicKeys.collection) >= 1
}

func (activityPubPublicKeys *activityPubPublicKeys) First() IActivityPubPublicKey {
	return activityPubPublicKeys.collection[0]
}

func (activityPubPublicKeys *activityPubPublicKeys) Append(activityPubPublicKey IActivityPubPublicKey) {
	activityPubPublicKeys.collection = append(activityPubPublicKeys.collection, activityPubPublicKey)
}

func (activityPubPublicKeys *activityPubPublicKeys) Reverse() IActivityPubPublicKeyCollection {
	slice := activityPubPublicKeys.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	activityPubPublicKeys.collection = slice

	return activityPubPublicKeys
}

func (activityPubPublicKeys *activityPubPublicKeys) ForEach(iterator ActivityPubPublicKeyIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubPublicKeys.collection {
		iterator(value)
	}
}

func (activityPubPublicKeys *activityPubPublicKeys) Array() ActivityPubPublicKeys {
	return activityPubPublicKeys.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubPublicKeyExists(id int64) bool {
	return dispatcher.conductor.ActivityPubPublicKeyManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubPublicKeyExistsWhich(condition ActivityPubPublicKeyCondition) bool {
	return dispatcher.conductor.ActivityPubPublicKeyManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubPublicKeys() IActivityPubPublicKeyCollection {
	return dispatcher.conductor.ActivityPubPublicKeyManager().ListActivityPubPublicKeys(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubPublicKey(iterator ActivityPubPublicKeyIterator) {
	dispatcher.conductor.ActivityPubPublicKeyManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubPublicKeys(predicate ActivityPubPublicKeyFilterPredicate) IActivityPubPublicKeyCollection {
	return dispatcher.conductor.ActivityPubPublicKeyManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubPublicKeys(predicate ActivityPubPublicKeyMapPredicate) IActivityPubPublicKeyCollection {
	return dispatcher.conductor.ActivityPubPublicKeyManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubPublicKey(id int64) IActivityPubPublicKey {
	if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().GetActivityPubPublicKey(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubPublicKey
	}
}

func (dispatcher *dispatcher) AddActivityPubPublicKey() IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddActivityPubPublicKeyAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddActivityPubPublicKey(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubPublicKeyWithCustomId(id int64) IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddActivityPubPublicKeyWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddActivityPubPublicKeyWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubPublicKey(source string, payload string) {
	dispatcher.conductor.ActivityPubPublicKeyManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubPublicKey(id int64) IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().UpdateActivityPubPublicKeyAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().UpdateActivityPubPublicKey(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubPublicKeyObject(object IObject, activityPubPublicKey IActivityPubPublicKey) IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().UpdateActivityPubPublicKeyAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().UpdateActivityPubPublicKey(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubPublicKeyObject(object IObject, activityPubPublicKey IActivityPubPublicKey) IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddOrUpdateActivityPubPublicKeyObjectAtomic(transaction, object.Id(), activityPubPublicKey, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().AddOrUpdateActivityPubPublicKeyObject(object.Id(), activityPubPublicKey, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubPublicKey(id int64) IActivityPubPublicKey {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().RemoveActivityPubPublicKeyAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	} else {
		if activityPubPublicKey, err := dispatcher.conductor.ActivityPubPublicKeyManager().RemoveActivityPubPublicKey(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubPublicKey
		}
	}
}
