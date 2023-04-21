package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type document struct {
	object
	content string
}

// noinspection GoUnusedExportedFunction
func InitializeDocument() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewDocument(id int64, content string) (IDocument, error) {
	instance := &document{
		object: object{
			id: id,
		},
		content: content,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewDocumentFromEntity(entity IDocumentEntity) (IDocument, error) {
	instance := &document{
		object: object{
			id: entity.Id(),
		},
		content: entity.Content(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (document *document) Content() string {
	return document.content
}

func (document *document) UpdateContent(content string, editor Identity) {
	if err := repository.Documents.UpdateContent(document.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}

	document.content = content
}

func (document *document) UpdateContentAtomic(transaction ITransaction, content string, editor Identity) {
	transaction.OnCommit(func() {
		document.content = content
	})

	if err := repository.Documents.UpdateContentAtomic(transaction, document.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (document *document) Validate() error {
	return nil
}

func (document *document) String() string {
	return fmt.Sprintf("Document (Id: %d, Content: %v)", document.Id(), document.Content())
}

//------------------------------------------------------------------------------

type documents struct {
	collection Documents
}

// NewDocuments creates an empty collection of 'Document' which is not thread-safe.
func NewDocuments() IDocumentCollection {
	return &documents{
		collection: make(Documents, 0),
	}
}

func (documents *documents) Count() int {
	return len(documents.collection)
}

func (documents *documents) IsEmpty() bool {
	return len(documents.collection) == 0
}

func (documents *documents) IsNotEmpty() bool {
	return len(documents.collection) > 0
}

func (documents *documents) HasExactlyOneItem() bool {
	return len(documents.collection) == 1
}

func (documents *documents) HasAtLeastOneItem() bool {
	return len(documents.collection) >= 1
}

func (documents *documents) First() IDocument {
	return documents.collection[0]
}

func (documents *documents) Append(document IDocument) {
	documents.collection = append(documents.collection, document)
}

func (documents *documents) ForEach(iterator DocumentIterator) {
	if iterator == nil {
		return
	}

	for _, value := range documents.collection {
		iterator(value)
	}
}

func (documents *documents) Array() Documents {
	return documents.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) DocumentExists(id int64) bool {
	return dispatcher.conductor.DocumentManager().Exists(id)
}

func (dispatcher *dispatcher) DocumentExistsWhich(condition DocumentCondition) bool {
	return dispatcher.conductor.DocumentManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListDocuments() IDocumentCollection {
	return dispatcher.conductor.DocumentManager().ListDocuments(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachDocument(iterator DocumentIterator) {
	dispatcher.conductor.DocumentManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterDocuments(predicate DocumentFilterPredicate) IDocumentCollection {
	return dispatcher.conductor.DocumentManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapDocuments(predicate DocumentMapPredicate) IDocumentCollection {
	return dispatcher.conductor.DocumentManager().Map(predicate)
}

func (dispatcher *dispatcher) GetDocument(id int64) IDocument {
	if document, err := dispatcher.conductor.DocumentManager().GetDocument(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return document
	}
}

func (dispatcher *dispatcher) AddDocument(content string) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().AddDocumentAtomic(transaction, content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().AddDocument(content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}

func (dispatcher *dispatcher) AddDocumentWithCustomId(id int64, content string) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().AddDocumentWithCustomIdAtomic(id, transaction, content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().AddDocumentWithCustomId(id, content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}

func (dispatcher *dispatcher) LogDocument(content string, source string, payload string) {
	dispatcher.conductor.DocumentManager().Log(content, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateDocument(id int64, content string) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().UpdateDocumentAtomic(transaction, id, content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().UpdateDocument(id, content, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateDocumentObject(object IObject, document IDocument) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().UpdateDocumentAtomic(transaction, object.Id(), document.Content(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().UpdateDocument(object.Id(), document.Content(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateDocumentObject(object IObject, document IDocument) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().AddOrUpdateDocumentObjectAtomic(transaction, object.Id(), document, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().AddOrUpdateDocumentObject(object.Id(), document, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}

func (dispatcher *dispatcher) RemoveDocument(id int64) IDocument {
	transaction := dispatcher.transaction
	if transaction != nil {
		if document, err := dispatcher.conductor.DocumentManager().RemoveDocumentAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	} else {
		if document, err := dispatcher.conductor.DocumentManager().RemoveDocument(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return document
		}
	}
}
