package model

type (
	ActivityPubFollowerEntities []IActivityPubFollowerEntity

	IActivityPubFollowerEntity interface {
		IEntity
		Handle() string
		Inbox() string
		Subject() string
		Activity() string
		Accepted() bool
	}

	IActivityPubFollowerPipeEntity interface {
		IActivityPubFollowerEntity
		IPipeEntity
	}

	IActivityPubFollowersRepository interface {
		IRepository
		Add(entity IActivityPubFollowerEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error
		FetchById(editor int64) (IActivityPubFollowerEntity, error)
		Update(entity IActivityPubFollowerEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error
		Remove(entity IActivityPubFollowerEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IActivityPubFollowerEntity, editor int64) error
		FetchAll() (ActivityPubFollowerEntities, error)
		UpdateHandle(id int64, value string, editor int64) error
		UpdateHandleAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateInbox(id int64, value string, editor int64) error
		UpdateInboxAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateSubject(id int64, value string, editor int64) error
		UpdateSubjectAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateActivity(id int64, value string, editor int64) error
		UpdateActivityAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateAccepted(id int64, value bool, editor int64) error
		UpdateAcceptedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
	}
)
