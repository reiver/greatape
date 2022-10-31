package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestRemoteActivitiesRepository_Add(test *testing.T) {
	type arguments struct {
		id            int64
		entryPoint    string
		duration      int64
		successful    bool
		errorMessage  string
		remoteAddress string
		userAgent     string
		eventType     uint32
		timestamp     int64
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
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewRemoteActivityEntity(testCase.arguments.id, testCase.arguments.entryPoint, testCase.arguments.duration, testCase.arguments.successful, testCase.arguments.errorMessage, testCase.arguments.remoteAddress, testCase.arguments.userAgent, testCase.arguments.eventType, testCase.arguments.timestamp)
			if result := RemoteActivities.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_FetchById(test *testing.T) {
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
			entity, err := RemoteActivities.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestRemoteActivitiesRepository_Update(test *testing.T) {
	type arguments struct {
		id            int64
		entryPoint    string
		duration      int64
		successful    bool
		errorMessage  string
		remoteAddress string
		userAgent     string
		eventType     uint32
		timestamp     int64
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
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:            0,
				entryPoint:    "entry_point",
				duration:      0,
				successful:    true,
				errorMessage:  "error_message",
				remoteAddress: "remote_address",
				userAgent:     "user_agent",
				eventType:     0,
				timestamp:     0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewRemoteActivityEntity(testCase.arguments.id, testCase.arguments.entryPoint, testCase.arguments.duration, testCase.arguments.successful, testCase.arguments.errorMessage, testCase.arguments.remoteAddress, testCase.arguments.userAgent, testCase.arguments.eventType, testCase.arguments.timestamp)
			if result := RemoteActivities.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_Remove(test *testing.T) {
	type arguments struct {
		id            int64
		entryPoint    string
		duration      int64
		successful    bool
		errorMessage  string
		remoteAddress string
		userAgent     string
		eventType     uint32
		timestamp     int64
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
			entity := NewRemoteActivityEntity(testCase.arguments.id, testCase.arguments.entryPoint, testCase.arguments.duration, testCase.arguments.successful, testCase.arguments.errorMessage, testCase.arguments.remoteAddress, testCase.arguments.userAgent, testCase.arguments.eventType, testCase.arguments.timestamp)
			if result := RemoteActivities.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_FetchAll(test *testing.T) {
	entities, err := RemoteActivities.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestRemoteActivitiesRepository_UpdateEntryPoint(test *testing.T) {
	type arguments struct {
		id         int64
		entryPoint string
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
				id:         0,
				entryPoint: "entry_point",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:         0,
				entryPoint: "entry_point",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:         0,
				entryPoint: "entry_point",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateEntryPoint(testCase.arguments.id, testCase.arguments.entryPoint, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateEntryPoint() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateDuration(test *testing.T) {
	type arguments struct {
		id       int64
		duration int64
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
				id:       0,
				duration: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				duration: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				duration: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateDuration(testCase.arguments.id, testCase.arguments.duration, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateDuration() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateSuccessful(test *testing.T) {
	type arguments struct {
		id         int64
		successful bool
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
				id:         0,
				successful: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:         0,
				successful: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:         0,
				successful: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateSuccessful(testCase.arguments.id, testCase.arguments.successful, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateSuccessful() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateErrorMessage(test *testing.T) {
	type arguments struct {
		id           int64
		errorMessage string
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
				id:           0,
				errorMessage: "error_message",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:           0,
				errorMessage: "error_message",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:           0,
				errorMessage: "error_message",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateErrorMessage(testCase.arguments.id, testCase.arguments.errorMessage, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateErrorMessage() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateRemoteAddress(test *testing.T) {
	type arguments struct {
		id            int64
		remoteAddress string
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
				id:            0,
				remoteAddress: "remote_address",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:            0,
				remoteAddress: "remote_address",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:            0,
				remoteAddress: "remote_address",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateRemoteAddress(testCase.arguments.id, testCase.arguments.remoteAddress, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateRemoteAddress() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateUserAgent(test *testing.T) {
	type arguments struct {
		id        int64
		userAgent string
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
				id:        0,
				userAgent: "user_agent",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				userAgent: "user_agent",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				userAgent: "user_agent",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateUserAgent(testCase.arguments.id, testCase.arguments.userAgent, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateUserAgent() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateEventType(test *testing.T) {
	type arguments struct {
		id        int64
		eventType uint32
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
				id:        0,
				eventType: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				eventType: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				eventType: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateEventType(testCase.arguments.id, testCase.arguments.eventType, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateEventType() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestRemoteActivitiesRepository_UpdateTimestamp(test *testing.T) {
	type arguments struct {
		id        int64
		timestamp int64
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
				id:        0,
				timestamp: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				timestamp: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				timestamp: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := RemoteActivities.UpdateTimestamp(testCase.arguments.id, testCase.arguments.timestamp, -1) == nil; result != testCase.expectation {
				test.Errorf("RemoteActivities.UpdateTimestamp() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
