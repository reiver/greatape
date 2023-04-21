package repository_test

import (
	"testing"

	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/reiver/greatape/components/model/repository"
)

func TestCategoriesRepository_Add(test *testing.T) {
	type arguments struct {
		id             int64
		categoryTypeId int64
		categoryId     int64
		title          string
		description    string
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
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewCategoryEntity(testCase.arguments.id, testCase.arguments.categoryTypeId, testCase.arguments.categoryId, testCase.arguments.title, testCase.arguments.description)
			if result := Categories.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Categories.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoriesRepository_FetchById(test *testing.T) {
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
			entity, err := Categories.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Categories.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestCategoriesRepository_Update(test *testing.T) {
	type arguments struct {
		id             int64
		categoryTypeId int64
		categoryId     int64
		title          string
		description    string
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
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:             0,
				categoryTypeId: 0,
				categoryId:     0,
				title:          "title",
				description:    "description",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewCategoryEntity(testCase.arguments.id, testCase.arguments.categoryTypeId, testCase.arguments.categoryId, testCase.arguments.title, testCase.arguments.description)
			if result := Categories.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Categories.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoriesRepository_Remove(test *testing.T) {
	type arguments struct {
		id             int64
		categoryTypeId int64
		categoryId     int64
		title          string
		description    string
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
			entity := NewCategoryEntity(testCase.arguments.id, testCase.arguments.categoryTypeId, testCase.arguments.categoryId, testCase.arguments.title, testCase.arguments.description)
			if result := Categories.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Categories.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoriesRepository_FetchAll(test *testing.T) {
	entities, err := Categories.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestCategoriesRepository_FetchAllByCategoryType(test *testing.T) {
	type arguments struct {
		categoryTypeId int64
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
				categoryTypeId: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				categoryTypeId: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				categoryTypeId: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entities, err := Categories.FetchAllByCategoryType(testCase.arguments.categoryTypeId)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Categories.FetchAllByCategoryType() = %v, expected %v", result, testCase.expectation)
			}

			_ = entities
		})
	}
}

func TestCategoriesRepository_FetchAllByCategory(test *testing.T) {
	type arguments struct {
		categoryId int64
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
				categoryId: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				categoryId: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				categoryId: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entities, err := Categories.FetchAllByCategory(testCase.arguments.categoryId)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Categories.FetchAllByCategory() = %v, expected %v", result, testCase.expectation)
			}

			_ = entities
		})
	}
}

func TestCategoriesRepository_UpdateTitle(test *testing.T) {
	type arguments struct {
		id    int64
		title string
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
				id:    0,
				title: "title",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				title: "title",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				title: "title",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Categories.UpdateTitle(testCase.arguments.id, testCase.arguments.title, -1) == nil; result != testCase.expectation {
				test.Errorf("Categories.UpdateTitle() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestCategoriesRepository_UpdateDescription(test *testing.T) {
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
			if result := Categories.UpdateDescription(testCase.arguments.id, testCase.arguments.description, -1) == nil; result != testCase.expectation {
				test.Errorf("Categories.UpdateDescription() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
