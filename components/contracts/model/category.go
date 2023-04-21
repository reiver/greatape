package model

type (
	CategoryEntities []ICategoryEntity

	ICategoryEntity interface {
		IEntity
		CategoryTypeId() int64
		CategoryId() int64
		Title() string
		Description() string
	}

	ICategoryPipeEntity interface {
		ICategoryEntity
		IPipeEntity
	}

	ICategoriesRepository interface {
		IRepository
		Add(entity ICategoryEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error
		FetchById(editor int64) (ICategoryEntity, error)
		Update(entity ICategoryEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error
		Remove(entity ICategoryEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity ICategoryEntity, editor int64) error
		FetchAll() (CategoryEntities, error)
		FetchAllByCategoryType(categoryTypeId int64) (CategoryEntities, error)
		FetchAllByCategory(categoryId int64) (CategoryEntities, error)
		UpdateTitle(id int64, value string, editor int64) error
		UpdateTitleAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateDescription(id int64, value string, editor int64) error
		UpdateDescriptionAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
