package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestCategoryTypesRepository_Add(test *testing.T) {
	type arguments struct {
		id          int64
		description string
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
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewCategoryTypeEntity(testCase.arguments.id, testCase.arguments.description)
			if result := CategoryTypes.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("CategoryTypes.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoryTypesRepository_FetchById(test *testing.T) {
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
			entity, err := CategoryTypes.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("CategoryTypes.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestCategoryTypesRepository_Update(test *testing.T) {
	type arguments struct {
		id          int64
		description string
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
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewCategoryTypeEntity(testCase.arguments.id, testCase.arguments.description)
			if result := CategoryTypes.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("CategoryTypes.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoryTypesRepository_Remove(test *testing.T) {
	type arguments struct {
		id          int64
		description string
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
			entity := NewCategoryTypeEntity(testCase.arguments.id, testCase.arguments.description)
			if result := CategoryTypes.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("CategoryTypes.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoryTypesRepository_FetchAll(test *testing.T) {
	entities, err := CategoryTypes.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestCategoryTypesRepository_UpdateDescription(test *testing.T) {
	type arguments struct {
		id          int64
		description string
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
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				description: "description",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := CategoryTypes.UpdateDescription(testCase.arguments.id, testCase.arguments.description, -1) == nil; result != testCase.expectation {
				test.Errorf("CategoryTypes.UpdateDescription() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
