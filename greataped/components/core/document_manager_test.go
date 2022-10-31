package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestDocumentManager_GetName(test *testing.T) {
	manager := Conductor.DocumentManager()

	if manager.Name() != DOCUMENT_MANAGER {
		test.Fail()
	}
}

func TestDocumentManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.DocumentManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestDocumentManager_Load(test *testing.T) {
	manager := Conductor.DocumentManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestDocumentManager_Reload(test *testing.T) {
	manager := Conductor.DocumentManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestDocumentManager_Count(test *testing.T) {
	manager := Conductor.DocumentManager()

	_ = manager.Count()
}

func TestDocumentManager_Exists(test *testing.T) {
	manager := Conductor.DocumentManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestDocumentManager_ListDocuments(test *testing.T) {
	manager := Conductor.DocumentManager()

	_ = manager.ListDocuments(0, 0, "", nil)
}

func TestDocumentManager_GetDocument(test *testing.T) {
	manager := Conductor.DocumentManager()

	if document, err := manager.GetDocument(0, nil); err == nil {
		_ = document
		test.FailNow()
	}
}

func TestDocumentManager_AddDocument(test *testing.T) {
	manager := Conductor.DocumentManager()

	document, err := manager.AddDocument("content", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = document
}

func TestDocumentManager_UpdateDocument(test *testing.T) {
	manager := Conductor.DocumentManager()

	document, err := manager.UpdateDocument(0, "content", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = document
}

func TestDocumentManager_RemoveDocument(test *testing.T) {
	manager := Conductor.DocumentManager()

	document, err := manager.RemoveDocument(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = document
}

func TestDocumentManager_Find(test *testing.T) {
	manager := Conductor.DocumentManager()

	document := manager.Find(0)
	if document == nil {
		test.Fail()
	}

	_ = document
}

func TestDocumentManager_ForEach(test *testing.T) {
	manager := Conductor.DocumentManager()

	manager.ForEach(func(document IDocument) {
		_ = document
	})
}

func TestDocumentManager_Filter(test *testing.T) {
	manager := Conductor.DocumentManager()

	documents := manager.Filter(func(document IDocument) bool {
		return document.Id() < 0
	})

	if documents.IsNotEmpty() {
		test.Fail()
	}

	_ = documents
}

func TestDocumentManager_Map(test *testing.T) {
	manager := Conductor.DocumentManager()

	documents := manager.Map(func(document IDocument) IDocument {
		return document
	})

	if documents.Count() != manager.Count() {
		test.Fail()
	}

	_ = documents
}
