package model

type (
	UserEntities []IUserEntity

	IUserEntity interface {
		IEntity
		Github() string
	}

	IUserPipeEntity interface {
		IUserEntity
		IPipeEntity
	}

	IUsersRepository interface {
		IRepository
		Add(entity IUserEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error
		FetchById(editor int64) (IUserEntity, error)
		Update(entity IUserEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error
		Remove(entity IUserEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IUserEntity, editor int64) error
		FetchAll() (UserEntities, error)
		UpdateGithub(id int64, value string, editor int64) error
		UpdateGithubAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
