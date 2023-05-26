package contracts

import . "github.com/xeronith/diamante/contracts/security"

var DocumentPassThroughFilter = func(IDocument) bool { return true }

type (
	Documents               []IDocument
	DocumentIterator        func(IDocument)
	DocumentCondition       func(IDocument) bool
	DocumentFilterPredicate func(IDocument) bool
	DocumentMapPredicate    func(IDocument) IDocument
	DocumentCacheCallback   func()

	IDocument interface {
		IObject
		// Content returns 'Content' of this 'Document' instance.
		Content() string
		// UpdateContent directly updates 'Content' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateContent(content string, editor Identity)
		// UpdateContentAtomic updates 'Content' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateContentAtomic(transaction ITransaction, content string, editor Identity)
	}

	IDocumentCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IDocument
		Append(document IDocument)
		ForEach(DocumentIterator)
		Reverse() IDocumentCollection
		Array() Documents
	}

	IDocumentManager interface {
		ISystemComponent
		OnCacheChanged(DocumentCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition DocumentCondition) bool
		ListDocuments(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IDocumentCollection
		GetDocument(id int64, editor Identity) (IDocument, error)
		AddDocument(content string, editor Identity) (IDocument, error)
		AddDocumentWithCustomId(id int64, content string, editor Identity) (IDocument, error)
		AddDocumentObject(document IDocument, editor Identity) (IDocument, error)
		AddDocumentAtomic(transaction ITransaction, content string, editor Identity) (IDocument, error)
		AddDocumentWithCustomIdAtomic(id int64, transaction ITransaction, content string, editor Identity) (IDocument, error)
		AddDocumentObjectAtomic(transaction ITransaction, document IDocument, editor Identity) (IDocument, error)
		Log(content string, source string, editor Identity, payload string)
		UpdateDocument(id int64, content string, editor Identity) (IDocument, error)
		UpdateDocumentObject(id int64, document IDocument, editor Identity) (IDocument, error)
		UpdateDocumentAtomic(transaction ITransaction, id int64, content string, editor Identity) (IDocument, error)
		UpdateDocumentObjectAtomic(transaction ITransaction, id int64, document IDocument, editor Identity) (IDocument, error)
		AddOrUpdateDocumentObject(id int64, document IDocument, editor Identity) (IDocument, error)
		AddOrUpdateDocumentObjectAtomic(transaction ITransaction, id int64, document IDocument, editor Identity) (IDocument, error)
		RemoveDocument(id int64, editor Identity) (IDocument, error)
		RemoveDocumentAtomic(transaction ITransaction, id int64, editor Identity) (IDocument, error)
		Find(id int64) IDocument
		ForEach(iterator DocumentIterator)
		Filter(predicate DocumentFilterPredicate) IDocumentCollection
		Map(predicate DocumentMapPredicate) IDocumentCollection
	}
)
