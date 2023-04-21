package repository_test

import (
	"testing"

	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/reiver/greatape/components/model/repository"
)

func TestActivityPubFollowersRepository_Add(test *testing.T) {
	type arguments struct {
		id       int64
		handle   string
		inbox    string
		subject  string
		activity string
		accepted bool
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
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewActivityPubFollowerEntity(testCase.arguments.id, testCase.arguments.handle, testCase.arguments.inbox, testCase.arguments.subject, testCase.arguments.activity, testCase.arguments.accepted)
			if result := ActivityPubFollowers.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_FetchById(test *testing.T) {
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
			entity, err := ActivityPubFollowers.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestActivityPubFollowersRepository_Update(test *testing.T) {
	type arguments struct {
		id       int64
		handle   string
		inbox    string
		subject  string
		activity string
		accepted bool
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
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				handle:   "handle",
				inbox:    "inbox",
				subject:  "subject",
				activity: "activity",
				accepted: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewActivityPubFollowerEntity(testCase.arguments.id, testCase.arguments.handle, testCase.arguments.inbox, testCase.arguments.subject, testCase.arguments.activity, testCase.arguments.accepted)
			if result := ActivityPubFollowers.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_Remove(test *testing.T) {
	type arguments struct {
		id       int64
		handle   string
		inbox    string
		subject  string
		activity string
		accepted bool
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
			entity := NewActivityPubFollowerEntity(testCase.arguments.id, testCase.arguments.handle, testCase.arguments.inbox, testCase.arguments.subject, testCase.arguments.activity, testCase.arguments.accepted)
			if result := ActivityPubFollowers.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_FetchAll(test *testing.T) {
	entities, err := ActivityPubFollowers.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestActivityPubFollowersRepository_UpdateHandle(test *testing.T) {
	type arguments struct {
		id     int64
		handle string
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
				handle: "handle",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				handle: "handle",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				handle: "handle",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubFollowers.UpdateHandle(testCase.arguments.id, testCase.arguments.handle, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.UpdateHandle() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_UpdateInbox(test *testing.T) {
	type arguments struct {
		id    int64
		inbox string
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
				inbox: "inbox",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				inbox: "inbox",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				inbox: "inbox",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubFollowers.UpdateInbox(testCase.arguments.id, testCase.arguments.inbox, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.UpdateInbox() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_UpdateSubject(test *testing.T) {
	type arguments struct {
		id      int64
		subject string
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
				subject: "subject",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				subject: "subject",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				subject: "subject",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubFollowers.UpdateSubject(testCase.arguments.id, testCase.arguments.subject, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.UpdateSubject() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_UpdateActivity(test *testing.T) {
	type arguments struct {
		id       int64
		activity string
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
				activity: "activity",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				activity: "activity",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				activity: "activity",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubFollowers.UpdateActivity(testCase.arguments.id, testCase.arguments.activity, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.UpdateActivity() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestActivityPubFollowersRepository_UpdateAccepted(test *testing.T) {
	type arguments struct {
		id       int64
		accepted bool
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
				accepted: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				accepted: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				accepted: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := ActivityPubFollowers.UpdateAccepted(testCase.arguments.id, testCase.arguments.accepted, -1) == nil; result != testCase.expectation {
				test.Errorf("ActivityPubFollowers.UpdateAccepted() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
