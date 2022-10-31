package contracts

import . "github.com/xeronith/diamante/contracts/security"

var CustomErrorPassThroughFilter = func(ICustomError) bool { return true }

type (
	CustomErrors               []ICustomError
	CustomErrorIterator        func(ICustomError)
	CustomErrorCondition       func(ICustomError) bool
	CustomErrorFilterPredicate func(ICustomError) bool
	CustomErrorMapPredicate    func(ICustomError) ICustomError
	CustomErrorCacheCallback   func()

	ICustomError interface {
	}

	ICustomErrorCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ICustomError
		Append(customError ICustomError)
		ForEach(CustomErrorIterator)
		Array() CustomErrors
	}

	ICustomErrorManager interface {
		ISystemComponent
		OnCacheChanged(CustomErrorCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition CustomErrorCondition) bool
		ListCustomErrors(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICustomErrorCollection
		GetCustomError(id int64, editor Identity) (ICustomError, error)
		AddCustomError(editor Identity) (ICustomError, error)
		AddCustomErrorWithCustomId(id int64, editor Identity) (ICustomError, error)
		AddCustomErrorObject(customError ICustomError, editor Identity) (ICustomError, error)
		AddCustomErrorAtomic(transaction ITransaction, editor Identity) (ICustomError, error)
		AddCustomErrorWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ICustomError, error)
		AddCustomErrorObjectAtomic(transaction ITransaction, customError ICustomError, editor Identity) (ICustomError, error)
		Log(source string, editor Identity, payload string)
		UpdateCustomError(id int64, editor Identity) (ICustomError, error)
		UpdateCustomErrorObject(id int64, customError ICustomError, editor Identity) (ICustomError, error)
		UpdateCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error)
		UpdateCustomErrorObjectAtomic(transaction ITransaction, id int64, customError ICustomError, editor Identity) (ICustomError, error)
		AddOrUpdateCustomErrorObject(id int64, customError ICustomError, editor Identity) (ICustomError, error)
		AddOrUpdateCustomErrorObjectAtomic(transaction ITransaction, id int64, customError ICustomError, editor Identity) (ICustomError, error)
		RemoveCustomError(id int64, editor Identity) (ICustomError, error)
		RemoveCustomErrorAtomic(transaction ITransaction, id int64, editor Identity) (ICustomError, error)
		Find(id int64) ICustomError
		ForEach(iterator CustomErrorIterator)
		Filter(predicate CustomErrorFilterPredicate) ICustomErrorCollection
		Map(predicate CustomErrorMapPredicate) ICustomErrorCollection
		ResolveError(document IDocument, editor Identity) (IResolveErrorResult, error)
	}

	IResolveErrorResult interface {
	}
)
