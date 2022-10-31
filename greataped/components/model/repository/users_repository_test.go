package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestUsersRepository_Add(test *testing.T) {
	type arguments struct {
		id     int64
		github string
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
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewUserEntity(testCase.arguments.id, testCase.arguments.github)
			if result := Users.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Users.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestUsersRepository_FetchById(test *testing.T) {
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
			entity, err := Users.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Users.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestUsersRepository_Update(test *testing.T) {
	type arguments struct {
		id     int64
		github string
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
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewUserEntity(testCase.arguments.id, testCase.arguments.github)
			if result := Users.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Users.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestUsersRepository_Remove(test *testing.T) {
	type arguments struct {
		id     int64
		github string
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
			entity := NewUserEntity(testCase.arguments.id, testCase.arguments.github)
			if result := Users.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Users.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestUsersRepository_FetchAll(test *testing.T) {
	entities, err := Users.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestUsersRepository_UpdateGithub(test *testing.T) {
	type arguments struct {
		id     int64
		github string
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
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				github: "github",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Users.UpdateGithub(testCase.arguments.id, testCase.arguments.github, -1) == nil; result != testCase.expectation {
				test.Errorf("Users.UpdateGithub() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
