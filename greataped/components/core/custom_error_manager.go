package core

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

// noinspection GoSnakeCaseUsage
const CUSTOM_ERROR_MANAGER = "CustomErrorManager"

type customErrorManager struct {
	systemComponent
	cache ICache
}

func newCustomErrorManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ICustomErrorManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &customErrorManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *customErrorManager) Name() string {
	return CUSTOM_ERROR_MANAGER
}

func (manager *customErrorManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *customErrorManager) Load() error {
	return nil
}

func (manager *customErrorManager) Reload() error {
	return manager.Load()
}

func (manager *customErrorManager) OnCacheChanged(callback CustomErrorCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *customErrorManager) Count() int {
	return manager.cache.Size()
}

func (manager *customErrorManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *customErrorManager) ExistsWhich(condition CustomErrorCondition) bool {
	var customErrors CustomErrors
	manager.ForEach(func(customError ICustomError) {
		if condition(customError) {
			customErrors = append(customErrors, customError)
		}
	})

	return len(customErrors) > 0
}

func (manager *customErrorManager) ListCustomErrors(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ICustomErrorCollection {
	return manager.Filter(CustomErrorPassThroughFilter)
}

func (manager *customErrorManager) GetCustomError(id int64, _ Identity) (ICustomError, error) {
	if customError := manager.Find(id); customError == nil {
		return nil, ERROR_CUSTOM_ERROR_NOT_FOUND
	} else {
		return customError, nil
	}
}

func (manager *customErrorManager) AddCustomError(editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddCustomErrorWithCustomId(id int64, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddCustomErrorObject(customError ICustomError, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddCustomErrorAtomic(transaction ITransaction, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddCustomErrorWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddCustomErrorObjectAtomic(transaction ITransaction, customError ICustomError, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) Log(source string, editor Identity, payload string) {
}

func (manager *customErrorManager) UpdateCustomError(id int64, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) UpdateCustomErrorObject(id int64, customError ICustomError, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) UpdateCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) UpdateCustomErrorObjectAtomic(transaction ITransaction, id int64, customError ICustomError, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) AddOrUpdateCustomErrorObject(id int64, customError ICustomError, editor Identity) (ICustomError, error) {
	if manager.Exists(id) {
		return manager.UpdateCustomErrorObject(id, customError, editor)
	} else {
		return manager.AddCustomErrorObject(customError, editor)
	}
}

func (manager *customErrorManager) AddOrUpdateCustomErrorObjectAtomic(transaction ITransaction, id int64, customError ICustomError, editor Identity) (ICustomError, error) {
	if manager.Exists(id) {
		return manager.UpdateCustomErrorObjectAtomic(transaction, id, customError, editor)
	} else {
		return manager.AddCustomErrorObjectAtomic(transaction, customError, editor)
	}
}

func (manager *customErrorManager) RemoveCustomError(id int64, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) RemoveCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *customErrorManager) Find(id int64) ICustomError {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ICustomError)
	}
}

func (manager *customErrorManager) ForEach(iterator CustomErrorIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ICustomError))
	})
}

func (manager *customErrorManager) Filter(predicate CustomErrorFilterPredicate) ICustomErrorCollection {
	customErrors := NewCustomErrors()
	if predicate == nil {
		return customErrors
	}

	manager.ForEach(func(customError ICustomError) {
		if predicate(customError) {
			customErrors.Append(customError)
		}
	})

	return customErrors
}

func (manager *customErrorManager) Map(predicate CustomErrorMapPredicate) ICustomErrorCollection {
	customErrors := NewCustomErrors()
	if predicate == nil {
		return customErrors
	}

	manager.ForEach(func(customError ICustomError) {
		customErrors.Append(predicate(customError))
	})

	return customErrors
}

//region IResolveErrorResult Implementation

type resolveErrorResult struct {
}

func NewResolveErrorResult(_ interface{}) IResolveErrorResult {
	return &resolveErrorResult{}
}

//endregion

func (manager *customErrorManager) ResolveError(document IDocument, editor Identity) (result IResolveErrorResult, err error) {
	return nil, ERROR_NOT_IMPLEMENTED
}
