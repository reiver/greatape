package core

import (
	"fmt"

	. "github.com/xeronith/diamante/contracts/security"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/model/repository"
)

type customError struct {
}

// noinspection GoUnusedExportedFunction
func InitializeCustomError() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewCustomError() (ICustomError, error) {
	instance := &customError{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (customError *customError) Validate() error {
	return nil
}

func (customError *customError) String() string {
	return fmt.Sprintf("CustomError (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type customErrors struct {
	collection CustomErrors
}

// NewCustomErrors creates an empty collection of 'Custom Error' which is not thread-safe.
func NewCustomErrors() ICustomErrorCollection {
	return &customErrors{
		collection: make(CustomErrors, 0),
	}
}

func (customErrors *customErrors) Count() int {
	return len(customErrors.collection)
}

func (customErrors *customErrors) IsEmpty() bool {
	return len(customErrors.collection) == 0
}

func (customErrors *customErrors) IsNotEmpty() bool {
	return len(customErrors.collection) > 0
}

func (customErrors *customErrors) HasExactlyOneItem() bool {
	return len(customErrors.collection) == 1
}

func (customErrors *customErrors) HasAtLeastOneItem() bool {
	return len(customErrors.collection) >= 1
}

func (customErrors *customErrors) First() ICustomError {
	return customErrors.collection[0]
}

func (customErrors *customErrors) Append(customError ICustomError) {
	customErrors.collection = append(customErrors.collection, customError)
}

func (customErrors *customErrors) ForEach(iterator CustomErrorIterator) {
	if iterator == nil {
		return
	}

	for _, value := range customErrors.collection {
		iterator(value)
	}
}

func (customErrors *customErrors) Array() CustomErrors {
	return customErrors.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) CustomErrorExists(id int64) bool {
	return dispatcher.conductor.CustomErrorManager().Exists(id)
}

func (dispatcher *dispatcher) CustomErrorExistsWhich(condition CustomErrorCondition) bool {
	return dispatcher.conductor.CustomErrorManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListCustomErrors() ICustomErrorCollection {
	return dispatcher.conductor.CustomErrorManager().ListCustomErrors(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachCustomError(iterator CustomErrorIterator) {
	dispatcher.conductor.CustomErrorManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterCustomErrors(predicate CustomErrorFilterPredicate) ICustomErrorCollection {
	return dispatcher.conductor.CustomErrorManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapCustomErrors(predicate CustomErrorMapPredicate) ICustomErrorCollection {
	return dispatcher.conductor.CustomErrorManager().Map(predicate)
}

func (dispatcher *dispatcher) GetCustomError(id int64) ICustomError {
	if customError, err := dispatcher.conductor.CustomErrorManager().GetCustomError(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return customError
	}
}

func (dispatcher *dispatcher) AddCustomError() ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddCustomErrorAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddCustomError(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

func (dispatcher *dispatcher) AddCustomErrorWithCustomId(id int64) ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddCustomErrorWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddCustomErrorWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

func (dispatcher *dispatcher) LogCustomError(source string, payload string) {
	dispatcher.conductor.CustomErrorManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateCustomError(id int64) ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().UpdateCustomErrorAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().UpdateCustomError(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateCustomErrorObject(object IObject, customError ICustomError) ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().UpdateCustomErrorAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().UpdateCustomError(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateCustomErrorObject(object IObject, customError ICustomError) ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddOrUpdateCustomErrorObjectAtomic(transaction, object.Id(), customError, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().AddOrUpdateCustomErrorObject(object.Id(), customError, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

func (dispatcher *dispatcher) RemoveCustomError(id int64) ICustomError {
	transaction := dispatcher.transaction
	if transaction != nil {
		if customError, err := dispatcher.conductor.CustomErrorManager().RemoveCustomErrorAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	} else {
		if customError, err := dispatcher.conductor.CustomErrorManager().RemoveCustomError(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return customError
		}
	}
}

func (dispatcher *dispatcher) ResolveError(document IDocument) (IResolveErrorResult, error) {
	return dispatcher.conductor.CustomErrorManager().ResolveError(document, dispatcher.identity)
}
