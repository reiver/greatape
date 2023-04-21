package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestActivityPubFollowerManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if manager.Name() != ACTIVITY_PUB_FOLLOWER_MANAGER {
		test.Fail()
	}
}

func TestActivityPubFollowerManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubFollowerManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubFollowerManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubFollowerManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	_ = manager.Count()
}

func TestActivityPubFollowerManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubFollowerManager_ListActivityPubFollowers(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	_ = manager.ListActivityPubFollowers(0, 0, "", nil)
}

func TestActivityPubFollowerManager_GetActivityPubFollower(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	if activityPubFollower, err := manager.GetActivityPubFollower(0, nil); err == nil {
		_ = activityPubFollower
		test.FailNow()
	}
}

func TestActivityPubFollowerManager_AddActivityPubFollower(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollower, err := manager.AddActivityPubFollower("handle", "inbox", "subject", "activity", true, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubFollower
}

func TestActivityPubFollowerManager_UpdateActivityPubFollower(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollower, err := manager.UpdateActivityPubFollower(0, "handle", "inbox", "subject", "activity", true, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubFollower
}

func TestActivityPubFollowerManager_RemoveActivityPubFollower(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollower, err := manager.RemoveActivityPubFollower(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubFollower
}

func TestActivityPubFollowerManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollower := manager.Find(0)
	if activityPubFollower == nil {
		test.Fail()
	}

	_ = activityPubFollower
}

func TestActivityPubFollowerManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	manager.ForEach(func(activityPubFollower IActivityPubFollower) {
		_ = activityPubFollower
	})
}

func TestActivityPubFollowerManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollowers := manager.Filter(func(activityPubFollower IActivityPubFollower) bool {
		return activityPubFollower.Id() < 0
	})

	if activityPubFollowers.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubFollowers
}

func TestActivityPubFollowerManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubFollowerManager()

	activityPubFollowers := manager.Map(func(activityPubFollower IActivityPubFollower) IActivityPubFollower {
		return activityPubFollower
	})

	if activityPubFollowers.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubFollowers
}
