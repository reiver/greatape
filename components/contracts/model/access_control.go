package model

type (
	AccessControlEntities []IAccessControlEntity

	IAccessControlEntity interface {
		IEntity
		Key() uint64
		Value() uint64
	}

	IAccessControlPipeEntity interface {
		IAccessControlEntity
		IPipeEntity
	}

	IAccessControlsRepository interface {
		IRepository
		Add(entity IAccessControlEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error
		FetchById(editor int64) (IAccessControlEntity, error)
		Update(entity IAccessControlEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error
		Remove(entity IAccessControlEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IAccessControlEntity, editor int64) error
		FetchAll() (AccessControlEntities, error)
		UpdateKey(id int64, value uint64, editor int64) error
		UpdateKeyAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error
		UpdateValue(id int64, value uint64, editor int64) error
		UpdateValueAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error
	}
)
