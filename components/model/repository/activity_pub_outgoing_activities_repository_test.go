package repository_test

import (
	"testing"

	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/reiver/greatape/components/model/repository"
)

func TestActivityPubOutgoingActivitiesRepository_Add(test *testing.T) {
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
			entity := NewActivityPubOutgoingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubOutgoingActivities.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_FetchById(test *testing.T) {
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
			entity, err := ActivityPubOutgoingActivities.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_Update(test *testing.T) {
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
			entity := NewActivityPubOutgoingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubOutgoingActivities.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_Remove(test *testing.T) {
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
			entity := NewActivityPubOutgoingActivityEntity(testCase.arguments.id, testCase.arguments.identityId, testCase.arguments.uniqueIdentifier, testCase.arguments.timestamp, testCase.arguments.from, testCase.arguments.to, testCase.arguments.content, testCase.arguments.raw)
			if result := ActivityPubOutgoingActivities.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_FetchAll(test *testing.T) {
	entities, err := ActivityPubOutgoingActivities.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestActivityPubOutgoingActivitiesRepository_FetchAllByIdentity(test *testing.T) {
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
			entities, err := ActivityPubOutgoingActivities.FetchAllByIdentity(testCase.arguments.identityId)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.FetchAllByIdentity() = %v, expected %v", result, testCase.expectation)
			}

			_ = entities
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateUniqueIdentifier(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateUniqueIdentifier(testCase.arguments.id, testCase.arguments.uniqueIdentifier, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateUniqueIdentifier() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateTimestamp(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateTimestamp(testCase.arguments.id, testCase.arguments.timestamp, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateTimestamp() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateFrom(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateFrom(testCase.arguments.id, testCase.arguments.from, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateFrom() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateTo(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateTo(testCase.arguments.id, testCase.arguments.to, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateTo() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateContent(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateContent(testCase.arguments.id, testCase.arguments.content, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateContent() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubOutgoingActivitiesRepository_UpdateRaw(test *testing.T) {
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
			if result := ActivityPubOutgoingActivities.UpdateRaw(testCase.arguments.id, testCase.arguments.raw, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubOutgoingActivities.UpdateRaw() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
