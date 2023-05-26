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

type remoteActivity struct {
	object
	entryPoint    string
	duration      int64
	successful    bool
	errorMessage  string
	remoteAddress string
	userAgent     string
	eventType     uint32
	timestamp     int64
}

// noinspection GoUnusedExportedFunction
func InitializeRemoteActivity() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error) {
	instance := &remoteActivity{
		object: object{
			id: id,
		},
		entryPoint:    entryPoint,
		duration:      duration,
		successful:    successful,
		errorMessage:  errorMessage,
		remoteAddress: remoteAddress,
		userAgent:     userAgent,
		eventType:     eventType,
		timestamp:     timestamp,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewRemoteActivityFromEntity(entity IRemoteActivityEntity) (IRemoteActivity, error) {
	instance := &remoteActivity{
		object: object{
			id: entity.Id(),
		},
		entryPoint:    entity.EntryPoint(),
		duration:      entity.Duration(),
		successful:    entity.Successful(),
		errorMessage:  entity.ErrorMessage(),
		remoteAddress: entity.RemoteAddress(),
		userAgent:     entity.UserAgent(),
		eventType:     entity.EventType(),
		timestamp:     entity.Timestamp(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (remoteActivity *remoteActivity) EntryPoint() string {
	return remoteActivity.entryPoint
}

func (remoteActivity *remoteActivity) UpdateEntryPoint(entryPoint string, editor Identity) {
	if err := repository.RemoteActivities.UpdateEntryPoint(remoteActivity.id, entryPoint, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.entryPoint = entryPoint
}

func (remoteActivity *remoteActivity) UpdateEntryPointAtomic(transaction ITransaction, entryPoint string, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.entryPoint = entryPoint
	})

	if err := repository.RemoteActivities.UpdateEntryPointAtomic(transaction, remoteActivity.id, entryPoint, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) Duration() int64 {
	return remoteActivity.duration
}

func (remoteActivity *remoteActivity) UpdateDuration(duration int64, editor Identity) {
	if err := repository.RemoteActivities.UpdateDuration(remoteActivity.id, duration, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.duration = duration
}

func (remoteActivity *remoteActivity) UpdateDurationAtomic(transaction ITransaction, duration int64, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.duration = duration
	})

	if err := repository.RemoteActivities.UpdateDurationAtomic(transaction, remoteActivity.id, duration, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) Successful() bool {
	return remoteActivity.successful
}

func (remoteActivity *remoteActivity) UpdateSuccessful(successful bool, editor Identity) {
	if err := repository.RemoteActivities.UpdateSuccessful(remoteActivity.id, successful, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.successful = successful
}

func (remoteActivity *remoteActivity) UpdateSuccessfulAtomic(transaction ITransaction, successful bool, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.successful = successful
	})

	if err := repository.RemoteActivities.UpdateSuccessfulAtomic(transaction, remoteActivity.id, successful, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) ErrorMessage() string {
	return remoteActivity.errorMessage
}

func (remoteActivity *remoteActivity) UpdateErrorMessage(errorMessage string, editor Identity) {
	if err := repository.RemoteActivities.UpdateErrorMessage(remoteActivity.id, errorMessage, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.errorMessage = errorMessage
}

func (remoteActivity *remoteActivity) UpdateErrorMessageAtomic(transaction ITransaction, errorMessage string, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.errorMessage = errorMessage
	})

	if err := repository.RemoteActivities.UpdateErrorMessageAtomic(transaction, remoteActivity.id, errorMessage, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) RemoteAddress() string {
	return remoteActivity.remoteAddress
}

func (remoteActivity *remoteActivity) UpdateRemoteAddress(remoteAddress string, editor Identity) {
	if err := repository.RemoteActivities.UpdateRemoteAddress(remoteActivity.id, remoteAddress, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.remoteAddress = remoteAddress
}

func (remoteActivity *remoteActivity) UpdateRemoteAddressAtomic(transaction ITransaction, remoteAddress string, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.remoteAddress = remoteAddress
	})

	if err := repository.RemoteActivities.UpdateRemoteAddressAtomic(transaction, remoteActivity.id, remoteAddress, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) UserAgent() string {
	return remoteActivity.userAgent
}

func (remoteActivity *remoteActivity) UpdateUserAgent(userAgent string, editor Identity) {
	if err := repository.RemoteActivities.UpdateUserAgent(remoteActivity.id, userAgent, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.userAgent = userAgent
}

func (remoteActivity *remoteActivity) UpdateUserAgentAtomic(transaction ITransaction, userAgent string, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.userAgent = userAgent
	})

	if err := repository.RemoteActivities.UpdateUserAgentAtomic(transaction, remoteActivity.id, userAgent, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) EventType() uint32 {
	return remoteActivity.eventType
}

func (remoteActivity *remoteActivity) UpdateEventType(eventType uint32, editor Identity) {
	if err := repository.RemoteActivities.UpdateEventType(remoteActivity.id, eventType, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.eventType = eventType
}

func (remoteActivity *remoteActivity) UpdateEventTypeAtomic(transaction ITransaction, eventType uint32, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.eventType = eventType
	})

	if err := repository.RemoteActivities.UpdateEventTypeAtomic(transaction, remoteActivity.id, eventType, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) Timestamp() int64 {
	return remoteActivity.timestamp
}

func (remoteActivity *remoteActivity) UpdateTimestamp(timestamp int64, editor Identity) {
	if err := repository.RemoteActivities.UpdateTimestamp(remoteActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}

	remoteActivity.timestamp = timestamp
}

func (remoteActivity *remoteActivity) UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity) {
	transaction.OnCommit(func() {
		remoteActivity.timestamp = timestamp
	})

	if err := repository.RemoteActivities.UpdateTimestampAtomic(transaction, remoteActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (remoteActivity *remoteActivity) Validate() error {
	return nil
}

func (remoteActivity *remoteActivity) String() string {
	return fmt.Sprintf("RemoteActivity (Id: %d, EntryPoint: %v, Duration: %v, Successful: %v, ErrorMessage: %v, RemoteAddress: %v, UserAgent: %v, EventType: %v, Timestamp: %v)", remoteActivity.Id(), remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp())
}

//------------------------------------------------------------------------------

type remoteActivities struct {
	collection RemoteActivities
}

// NewRemoteActivities creates an empty collection of 'Remote Activity' which is not thread-safe.
func NewRemoteActivities() IRemoteActivityCollection {
	return &remoteActivities{
		collection: make(RemoteActivities, 0),
	}
}

func (remoteActivities *remoteActivities) Count() int {
	return len(remoteActivities.collection)
}

func (remoteActivities *remoteActivities) IsEmpty() bool {
	return len(remoteActivities.collection) == 0
}

func (remoteActivities *remoteActivities) IsNotEmpty() bool {
	return len(remoteActivities.collection) > 0
}

func (remoteActivities *remoteActivities) HasExactlyOneItem() bool {
	return len(remoteActivities.collection) == 1
}

func (remoteActivities *remoteActivities) HasAtLeastOneItem() bool {
	return len(remoteActivities.collection) >= 1
}

func (remoteActivities *remoteActivities) First() IRemoteActivity {
	return remoteActivities.collection[0]
}

func (remoteActivities *remoteActivities) Append(remoteActivity IRemoteActivity) {
	remoteActivities.collection = append(remoteActivities.collection, remoteActivity)
}

func (remoteActivities *remoteActivities) Reverse() IRemoteActivityCollection {
	slice := remoteActivities.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	remoteActivities.collection = slice

	return remoteActivities
}

func (remoteActivities *remoteActivities) ForEach(iterator RemoteActivityIterator) {
	if iterator == nil {
		return
	}

	for _, value := range remoteActivities.collection {
		iterator(value)
	}
}

func (remoteActivities *remoteActivities) Array() RemoteActivities {
	return remoteActivities.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) RemoteActivityExists(id int64) bool {
	return dispatcher.conductor.RemoteActivityManager().Exists(id)
}

func (dispatcher *dispatcher) RemoteActivityExistsWhich(condition RemoteActivityCondition) bool {
	return dispatcher.conductor.RemoteActivityManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListRemoteActivities() IRemoteActivityCollection {
	return dispatcher.conductor.RemoteActivityManager().ListRemoteActivities(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachRemoteActivity(iterator RemoteActivityIterator) {
	dispatcher.conductor.RemoteActivityManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterRemoteActivities(predicate RemoteActivityFilterPredicate) IRemoteActivityCollection {
	return dispatcher.conductor.RemoteActivityManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapRemoteActivities(predicate RemoteActivityMapPredicate) IRemoteActivityCollection {
	return dispatcher.conductor.RemoteActivityManager().Map(predicate)
}

func (dispatcher *dispatcher) GetRemoteActivity(id int64) IRemoteActivity {
	if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().GetRemoteActivity(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return remoteActivity
	}
}

func (dispatcher *dispatcher) AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddRemoteActivityAtomic(transaction, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddRemoteActivity(entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}

func (dispatcher *dispatcher) AddRemoteActivityWithCustomId(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddRemoteActivityWithCustomIdAtomic(id, transaction, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddRemoteActivityWithCustomId(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}

func (dispatcher *dispatcher) LogRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, payload string) {
	dispatcher.conductor.RemoteActivityManager().Log(entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().UpdateRemoteActivityAtomic(transaction, id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().UpdateRemoteActivity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateRemoteActivityObject(object IObject, remoteActivity IRemoteActivity) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().UpdateRemoteActivityAtomic(transaction, object.Id(), remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().UpdateRemoteActivity(object.Id(), remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateRemoteActivityObject(object IObject, remoteActivity IRemoteActivity) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddOrUpdateRemoteActivityObjectAtomic(transaction, object.Id(), remoteActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().AddOrUpdateRemoteActivityObject(object.Id(), remoteActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}

func (dispatcher *dispatcher) RemoveRemoteActivity(id int64) IRemoteActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().RemoveRemoteActivityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	} else {
		if remoteActivity, err := dispatcher.conductor.RemoteActivityManager().RemoveRemoteActivity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return remoteActivity
		}
	}
}
