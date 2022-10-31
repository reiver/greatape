package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestSystemSchedulesRepository_Add(test *testing.T) {
	type arguments struct {
		id      int64
		enabled bool
		config  string
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
				enabled: true,
				config:  "config",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
				config:  "config",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
				config:  "config",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewSystemScheduleEntity(testCase.arguments.id, testCase.arguments.enabled, testCase.arguments.config)
			if result := SystemSchedules.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestSystemSchedulesRepository_FetchById(test *testing.T) {
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
			entity, err := SystemSchedules.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestSystemSchedulesRepository_Update(test *testing.T) {
	type arguments struct {
		id      int64
		enabled bool
		config  string
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
				enabled: true,
				config:  "config",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
				config:  "config",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
				config:  "config",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewSystemScheduleEntity(testCase.arguments.id, testCase.arguments.enabled, testCase.arguments.config)
			if result := SystemSchedules.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestSystemSchedulesRepository_Remove(test *testing.T) {
	type arguments struct {
		id      int64
		enabled bool
		config  string
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
			entity := NewSystemScheduleEntity(testCase.arguments.id, testCase.arguments.enabled, testCase.arguments.config)
			if result := SystemSchedules.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestSystemSchedulesRepository_FetchAll(test *testing.T) {
	entities, err := SystemSchedules.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestSystemSchedulesRepository_UpdateEnabled(test *testing.T) {
	type arguments struct {
		id      int64
		enabled bool
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
				enabled: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				enabled: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := SystemSchedules.UpdateEnabled(testCase.arguments.id, testCase.arguments.enabled, -1) == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.UpdateEnabled() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestSystemSchedulesRepository_UpdateConfig(test *testing.T) {
	type arguments struct {
		id     int64
		config string
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
				config: "config",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				config: "config",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				config: "config",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := SystemSchedules.UpdateConfig(testCase.arguments.id, testCase.arguments.config, -1) == nil; result != testCase.expectation {
				test.Errorf("SystemSchedules.UpdateConfig() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
