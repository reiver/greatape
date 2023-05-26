package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubPublicKeyPassThroughFilter = func(IActivityPubPublicKey) bool { return true }

type (
	ActivityPubPublicKeys               []IActivityPubPublicKey
	ActivityPubPublicKeyIterator        func(IActivityPubPublicKey)
	ActivityPubPublicKeyCondition       func(IActivityPubPublicKey) bool
	ActivityPubPublicKeyFilterPredicate func(IActivityPubPublicKey) bool
	ActivityPubPublicKeyMapPredicate    func(IActivityPubPublicKey) IActivityPubPublicKey
	ActivityPubPublicKeyCacheCallback   func()

	IActivityPubPublicKey interface {
		// Id returns 'Id' of this 'ActivityPubPublicKey' instance.
		Id() string
		// SetId sets 'Id' in-memory value of this 'ActivityPubPublicKey' instance.
		// This doesn't affect the persistent data store.
		SetId(id string)
		// Owner returns 'Owner' of this 'ActivityPubPublicKey' instance.
		Owner() string
		// SetOwner sets 'Owner' in-memory value of this 'ActivityPubPublicKey' instance.
		// This doesn't affect the persistent data store.
		SetOwner(owner string)
		// PublicKeyPem returns 'PublicKeyPem' of this 'ActivityPubPublicKey' instance.
		PublicKeyPem() string
		// SetPublicKeyPem sets 'PublicKeyPem' in-memory value of this 'ActivityPubPublicKey' instance.
		// This doesn't affect the persistent data store.
		SetPublicKeyPem(publicKeyPem string)
	}

	IActivityPubPublicKeyCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubPublicKey
		Append(activityPubPublicKey IActivityPubPublicKey)
		ForEach(ActivityPubPublicKeyIterator)
		Reverse() IActivityPubPublicKeyCollection
		Array() ActivityPubPublicKeys
	}

	IActivityPubPublicKeyManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubPublicKeyCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubPublicKeyCondition) bool
		ListActivityPubPublicKeys(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubPublicKeyCollection
		GetActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKey(editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyWithCustomId(id int64, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyObject(activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyAtomic(transaction ITransaction, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyObjectAtomic(transaction ITransaction, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		Log(source string, editor Identity, payload string)
		UpdateActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		UpdateActivityPubPublicKeyObject(id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		UpdateActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error)
		UpdateActivityPubPublicKeyObjectAtomic(transaction ITransaction, id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		AddOrUpdateActivityPubPublicKeyObject(id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		AddOrUpdateActivityPubPublicKeyObjectAtomic(transaction ITransaction, id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error)
		RemoveActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		RemoveActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error)
		Find(id int64) IActivityPubPublicKey
		ForEach(iterator ActivityPubPublicKeyIterator)
		Filter(predicate ActivityPubPublicKeyFilterPredicate) IActivityPubPublicKeyCollection
		Map(predicate ActivityPubPublicKeyMapPredicate) IActivityPubPublicKeyCollection
	}
)
