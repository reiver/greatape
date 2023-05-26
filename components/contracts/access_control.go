package contracts

import . "github.com/xeronith/diamante/contracts/security"

var AccessControlPassThroughFilter = func(IAccessControl) bool { return true }

type (
	AccessControls               []IAccessControl
	AccessControlIterator        func(IAccessControl)
	AccessControlCondition       func(IAccessControl) bool
	AccessControlFilterPredicate func(IAccessControl) bool
	AccessControlMapPredicate    func(IAccessControl) IAccessControl
	AccessControlCacheCallback   func()

	IAccessControl interface {
		IObject
		// Key returns 'Key' of this 'AccessControl' instance.
		Key() uint64
		// UpdateKey directly updates 'Key' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateKey(key uint64, editor Identity)
		// UpdateKeyAtomic updates 'Key' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateKeyAtomic(transaction ITransaction, key uint64, editor Identity)
		// Value returns 'Value' of this 'AccessControl' instance.
		Value() uint64
		// UpdateValue directly updates 'Value' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateValue(value uint64, editor Identity)
		// UpdateValueAtomic updates 'Value' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateValueAtomic(transaction ITransaction, value uint64, editor Identity)
	}

	IAccessControlCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IAccessControl
		Append(accessControl IAccessControl)
		ForEach(AccessControlIterator)
		Reverse() IAccessControlCollection
		Array() AccessControls
	}

	IAccessControlManager interface {
		ISystemComponent
		OnCacheChanged(AccessControlCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition AccessControlCondition) bool
		ListAccessControls(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IAccessControlCollection
		GetAccessControl(id int64, editor Identity) (IAccessControl, error)
		AddOrUpdateAccessControl(key uint64, value uint64, editor Identity) error
		AccessControls() map[uint64]uint64
		AddAccessControl(key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlWithCustomId(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlObject(accessControl IAccessControl, editor Identity) (IAccessControl, error)
		AddAccessControlAtomic(transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlWithCustomIdAtomic(id int64, transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlObjectAtomic(transaction ITransaction, accessControl IAccessControl, editor Identity) (IAccessControl, error)
		Log(key uint64, value uint64, source string, editor Identity, payload string)
		UpdateAccessControl(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		UpdateAccessControlObject(id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error)
		UpdateAccessControlAtomic(transaction ITransaction, id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		UpdateAccessControlObjectAtomic(transaction ITransaction, id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error)
		AddOrUpdateAccessControlObject(id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error)
		AddOrUpdateAccessControlObjectAtomic(transaction ITransaction, id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error)
		RemoveAccessControl(id int64, editor Identity) (IAccessControl, error)
		RemoveAccessControlAtomic(transaction ITransaction, id int64, editor Identity) (IAccessControl, error)
		Find(id int64) IAccessControl
		ForEach(iterator AccessControlIterator)
		Filter(predicate AccessControlFilterPredicate) IAccessControlCollection
		Map(predicate AccessControlMapPredicate) IAccessControlCollection
	}
)
