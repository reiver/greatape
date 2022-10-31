package core

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
	"rail.town/infrastructure/components/model/repository"
)

// noinspection GoSnakeCaseUsage
const DOCUMENT_MANAGER = "DocumentManager"

type documentManager struct {
	systemComponent
	cache ICache
}

func newDocumentManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IDocumentManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &documentManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *documentManager) Name() string {
	return DOCUMENT_MANAGER
}

func (manager *documentManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *documentManager) Load() error {
	documentEntities, err := repository.Documents.FetchAll()
	if err != nil {
		return err
	}

	documents := make(SystemObjectCache)
	for _, documentEntity := range documentEntities {
		if document, err := NewDocumentFromEntity(documentEntity); err == nil {
			documents[document.Id()] = document
		} else {
			return err
		}
	}

	manager.cache.Load(documents)
	return nil
}

func (manager *documentManager) Reload() error {
	return manager.Load()
}

func (manager *documentManager) OnCacheChanged(callback DocumentCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *documentManager) Count() int {
	return manager.cache.Size()
}

func (manager *documentManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *documentManager) ExistsWhich(condition DocumentCondition) bool {
	var documents Documents
	manager.ForEach(func(document IDocument) {
		if condition(document) {
			documents = append(documents, document)
		}
	})

	return len(documents) > 0
}

func (manager *documentManager) ListDocuments(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IDocumentCollection {
	return manager.Filter(DocumentPassThroughFilter)
}

func (manager *documentManager) GetDocument(id int64, _ Identity) (IDocument, error) {
	if document := manager.Find(id); document == nil {
		return nil, ERROR_DOCUMENT_NOT_FOUND
	} else {
		return document, nil
	}
}

func (manager *documentManager) AddDocument(content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(manager.UniqueId(), content)
	return manager.Apply(documentEntity, repository.Documents.Add, manager.cache.Put, editor)
}

func (manager *documentManager) AddDocumentWithCustomId(id int64, content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, content)
	return manager.Apply(documentEntity, repository.Documents.Add, manager.cache.Put, editor)
}

func (manager *documentManager) AddDocumentObject(document IDocument, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(manager.UniqueId(), document.Content())
	return manager.Apply(documentEntity, repository.Documents.Add, manager.cache.Put, editor)
}

func (manager *documentManager) AddDocumentAtomic(transaction ITransaction, content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(manager.UniqueId(), content)
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.AddAtomic, manager.cache.Put, editor)
}

func (manager *documentManager) AddDocumentWithCustomIdAtomic(id int64, transaction ITransaction, content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, content)
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.AddAtomic, manager.cache.Put, editor)
}

func (manager *documentManager) AddDocumentObjectAtomic(transaction ITransaction, document IDocument, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(manager.UniqueId(), document.Content())
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.AddAtomic, manager.cache.Put, editor)
}

func (manager *documentManager) Log(content string, source string, editor Identity, payload string) {
	documentPipeEntity := NewDocumentPipeEntity(manager.UniqueId(), content, source, editor.Id(), payload)
	repository.Pipe.Insert(documentPipeEntity)

	document, err := NewDocumentFromEntity(documentPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(document.Id(), document)
	}
}

func (manager *documentManager) UpdateDocument(id int64, content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, content)
	return manager.Apply(documentEntity, repository.Documents.Update, manager.cache.Put, editor)
}

func (manager *documentManager) UpdateDocumentObject(id int64, document IDocument, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, document.Content())
	return manager.Apply(documentEntity, repository.Documents.Update, manager.cache.Put, editor)
}

func (manager *documentManager) UpdateDocumentAtomic(transaction ITransaction, id int64, content string, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, content)
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *documentManager) UpdateDocumentObjectAtomic(transaction ITransaction, id int64, document IDocument, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, document.Content())
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *documentManager) AddOrUpdateDocumentObject(id int64, document IDocument, editor Identity) (IDocument, error) {
	if manager.Exists(id) {
		return manager.UpdateDocumentObject(id, document, editor)
	} else {
		return manager.AddDocumentObject(document, editor)
	}
}

func (manager *documentManager) AddOrUpdateDocumentObjectAtomic(transaction ITransaction, id int64, document IDocument, editor Identity) (IDocument, error) {
	if manager.Exists(id) {
		return manager.UpdateDocumentObjectAtomic(transaction, id, document, editor)
	} else {
		return manager.AddDocumentObjectAtomic(transaction, document, editor)
	}
}

func (manager *documentManager) RemoveDocument(id int64, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, "")
	return manager.Apply(documentEntity, repository.Documents.Remove, manager.cache.Remove, editor)
}

func (manager *documentManager) RemoveDocumentAtomic(transaction ITransaction, id int64, editor Identity) (IDocument, error) {
	documentEntity := NewDocumentEntity(id, "")
	return manager.ApplyAtomic(transaction, documentEntity, repository.Documents.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *documentManager) Apply(documentEntity IDocumentEntity, repositoryHandler func(IDocumentEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IDocument, error) {
	result, err := NewDocumentFromEntity(documentEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(documentEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *documentManager) ApplyAtomic(transaction ITransaction, documentEntity IDocumentEntity, repositoryHandler func(IRepositoryTransaction, IDocumentEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IDocument, error) {
	result, err := NewDocumentFromEntity(documentEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, documentEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *documentManager) Find(id int64) IDocument {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IDocument)
	}
}

func (manager *documentManager) ForEach(iterator DocumentIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IDocument))
	})
}

func (manager *documentManager) Filter(predicate DocumentFilterPredicate) IDocumentCollection {
	documents := NewDocuments()
	if predicate == nil {
		return documents
	}

	manager.ForEach(func(document IDocument) {
		if predicate(document) {
			documents.Append(document)
		}
	})

	return documents
}

func (manager *documentManager) Map(predicate DocumentMapPredicate) IDocumentCollection {
	documents := NewDocuments()
	if predicate == nil {
		return documents
	}

	manager.ForEach(func(document IDocument) {
		documents.Append(predicate(document))
	})

	return documents
}
