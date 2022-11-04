package model

type (
	ActivityPubOutgoingActivityEntities []IActivityPubOutgoingActivityEntity

	IActivityPubOutgoingActivityEntity interface {
		IEntity
		IdentityId() int64
		UniqueIdentifier() string
		Timestamp() int64
		From() string
		To() string
		Content() string
		Raw() string
	}

	IActivityPubOutgoingActivityPipeEntity interface {
		IActivityPubOutgoingActivityEntity
		IPipeEntity
	}

	IActivityPubOutgoingActivitiesRepository interface {
		IRepository
		Add(entity IActivityPubOutgoingActivityEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error
		FetchById(editor int64) (IActivityPubOutgoingActivityEntity, error)
		Update(entity IActivityPubOutgoingActivityEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error
		Remove(entity IActivityPubOutgoingActivityEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IActivityPubOutgoingActivityEntity, editor int64) error
		FetchAll() (ActivityPubOutgoingActivityEntities, error)
		FetchAllByIdentity(identityId int64) (ActivityPubOutgoingActivityEntities, error)
		UpdateUniqueIdentifier(id int64, value string, editor int64) error
		UpdateUniqueIdentifierAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateTimestamp(id int64, value int64, editor int64) error
		UpdateTimestampAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error
		UpdateFrom(id int64, value string, editor int64) error
		UpdateFromAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateTo(id int64, value string, editor int64) error
		UpdateToAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateContent(id int64, value string, editor int64) error
		UpdateContentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateRaw(id int64, value string, editor int64) error
		UpdateRawAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
