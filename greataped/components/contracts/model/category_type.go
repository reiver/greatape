package model

type (
	CategoryTypeEntities []ICategoryTypeEntity

	ICategoryTypeEntity interface {
		IEntity
		Description() string
	}

	ICategoryTypePipeEntity interface {
		ICategoryTypeEntity
		IPipeEntity
	}

	ICategoryTypesRepository interface {
		IRepository
		Add(entity ICategoryTypeEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error
		FetchById(editor int64) (ICategoryTypeEntity, error)
		Update(entity ICategoryTypeEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error
		Remove(entity ICategoryTypeEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity ICategoryTypeEntity, editor int64) error
		FetchAll() (CategoryTypeEntities, error)
		UpdateDescription(id int64, value string, editor int64) error
		UpdateDescriptionAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
