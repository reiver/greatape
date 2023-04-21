package model

type (
	DocumentEntities []IDocumentEntity

	IDocumentEntity interface {
		IEntity
		Content() string
	}

	IDocumentPipeEntity interface {
		IDocumentEntity
		IPipeEntity
	}

	IDocumentsRepository interface {
		IRepository
		Add(entity IDocumentEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error
		FetchById(editor int64) (IDocumentEntity, error)
		Update(entity IDocumentEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error
		Remove(entity IDocumentEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IDocumentEntity, editor int64) error
		FetchAll() (DocumentEntities, error)
		UpdateContent(id int64, value string, editor int64) error
		UpdateContentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
	}
)
