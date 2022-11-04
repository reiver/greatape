package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestActivityPubIncomingActivitiesRepository_Add(test *testing.T) {
	type arguments struct {
		id               int64
		identityId       int64
		uniqueIdentifier string
		timestamp        int64
		from             string
		to               string
		content          string
		raw              string
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
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewActivityPubIncomingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubIncomingActivities.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_FetchById(test *testing.T) {
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
			entity, err := ActivityPubIncomingActivities.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_Update(test *testing.T) {
	type arguments struct {
		id               int64
		identityId       int64
		uniqueIdentifier string
		timestamp        int64
		from             string
		to               string
		content          string
		raw              string
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
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:               0,
				identityId:       0,
				uniqueIdentifier: "unique_identifier",
				timestamp:        0,
				from:             "from",
				to:               "to",
				content:          "content",
				raw:              "raw",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewActivityPubIncomingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubIncomingActivities.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_Remove(test *testing.T) {
	type arguments struct {
		id               int64
		identityId       int64
		uniqueIdentifier string
		timestamp        int64
		from             string
		to               string
		content          string
		raw              string
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
			entity := NewActivityPubIncomingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubIncomingActivities.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_FetchAll(test *testing.T) {
	entities, err := ActivityPubIncomingActivities.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestActivityPubIncomingActivitiesRepository_FetchAllByIdentity(test *testing.T) {
	type arguments struct {
		identityId int64
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
				identityId: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				identityId: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				identityId: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entities, err := ActivityPubIncomingActivities.FetchAllByIdentity(testCase.arguments.identityId)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.FetchAllByIdentity() = %v, expected %v", result, testCase.expectation)
			}

			_ = entities
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateUniqueIdentifier(test *testing.T) {
	type arguments struct {
		id               int64
		uniqueIdentifier string
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
				id:               0,
				uniqueIdentifier: "unique_identifier",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:               0,
				uniqueIdentifier: "unique_identifier",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:               0,
				uniqueIdentifier: "unique_identifier",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubIncomingActivities.UpdateUniqueIdentifier(testCase.arguments.id, testCase.arguments.uniqueIdentifier, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateUniqueIdentifier() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateTimestamp(test *testing.T) {
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
			if result := ActivityPubIncomingActivities.UpdateTimestamp(testCase.arguments.id, testCase.arguments.timestamp, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateTimestamp() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateFrom(test *testing.T) {
	type arguments struct {
		id   int64
		from string
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
				id:   0,
				from: "from",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:   0,
				from: "from",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:   0,
				from: "from",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubIncomingActivities.UpdateFrom(testCase.arguments.id, testCase.arguments.from, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateFrom() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateTo(test *testing.T) {
	type arguments struct {
		id int64
		to string
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
				to: "to",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id: 0,
				to: "to",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id: 0,
				to: "to",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubIncomingActivities.UpdateTo(testCase.arguments.id, testCase.arguments.to, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateTo() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateContent(test *testing.T) {
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
			if result := ActivityPubIncomingActivities.UpdateContent(testCase.arguments.id, testCase.arguments.content, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateContent() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubIncomingActivitiesRepository_UpdateRaw(test *testing.T) {
	type arguments struct {
		id  int64
		raw string
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
				raw: "raw",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:  0,
				raw: "raw",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:  0,
				raw: "raw",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubIncomingActivities.UpdateRaw(testCase.arguments.id, testCase.arguments.raw, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubIncomingActivities.UpdateRaw() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
