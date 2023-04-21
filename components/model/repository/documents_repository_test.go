package repository_test

import (
	"testing"

	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/reiver/greatape/components/model/repository"
)

func TestDocumentsRepository_Add(test *testing.T) {
	type arguments struct {
		id      int64
		content string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewDocumentEntity(testCase.arguments.id, testCase.arguments.content)
			if result := Documents.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Documents.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestDocumentsRepository_FetchById(test *testing.T) {
	type arguments struct {
		id int64
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity, err := Documents.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Documents.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestDocumentsRepository_Update(test *testing.T) {
	type arguments struct {
		id      int64
		content string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewDocumentEntity(testCase.arguments.id, testCase.arguments.content)
			if result := Documents.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Documents.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestDocumentsRepository_Remove(test *testing.T) {
	type arguments struct {
		id      int64
		content string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewDocumentEntity(testCase.arguments.id, testCase.arguments.content)
			if result := Documents.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Documents.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestDocumentsRepository_FetchAll(test *testing.T) {
	entities, err := Documents.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestDocumentsRepository_UpdateContent(test *testing.T) {
	type arguments struct {
		id      int64
		content string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				content: "content",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Documents.UpdateContent(testCase.arguments.id, testCase.arguments.content, -1) == nil; result != testCase.expectation {
				test.Errorf("Documents.UpdateContent() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
