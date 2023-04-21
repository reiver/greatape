package repository_test

import (
	"testing"

	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/reiver/greatape/components/model/repository"
)

func TestAccessControlsRepository_Add(test *testing.T) {
	type arguments struct {
		id    int64
		key   uint64
		value uint64
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
				key:   0,
				value: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				key:   0,
				value: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				key:   0,
				value: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewAccessControlEntity(testCase.arguments.id, testCase.arguments.key, testCase.arguments.value)
			if result := AccessControls.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("AccessControls.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestAccessControlsRepository_FetchById(test *testing.T) {
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
			entity, err := AccessControls.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("AccessControls.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestAccessControlsRepository_Update(test *testing.T) {
	type arguments struct {
		id    int64
		key   uint64
		value uint64
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
				key:   0,
				value: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				key:   0,
				value: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				key:   0,
				value: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewAccessControlEntity(testCase.arguments.id, testCase.arguments.key, testCase.arguments.value)
			if result := AccessControls.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("AccessControls.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestAccessControlsRepository_Remove(test *testing.T) {
	type arguments struct {
		id    int64
		key   uint64
		value uint64
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
			entity := NewAccessControlEntity(testCase.arguments.id, testCase.arguments.key, testCase.arguments.value)
			if result := AccessControls.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("AccessControls.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestAccessControlsRepository_FetchAll(test *testing.T) {
	entities, err := AccessControls.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestAccessControlsRepository_UpdateKey(test *testing.T) {
	type arguments struct {
		id  int64
		key uint64
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
				id:  0,
				key: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:  0,
				key: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:  0,
				key: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := AccessControls.UpdateKey(testCase.arguments.id, testCase.arguments.key, -1) == nil; result != testCase.expectation {
				test.Errorf("AccessControls.UpdateKey() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestAccessControlsRepository_UpdateValue(test *testing.T) {
	type arguments struct {
		id    int64
		value uint64
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
				value: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				value: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				value: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := AccessControls.UpdateValue(testCase.arguments.id, testCase.arguments.value, -1) == nil; result != testCase.expectation {
				test.Errorf("AccessControls.UpdateValue() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
