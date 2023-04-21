package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestActivityPubActivityManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if manager.Name() != ACTIVITY_PUB_ACTIVITY_MANAGER {
		test.Fail()
	}
}

func TestActivityPubActivityManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubActivityManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubActivityManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubActivityManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	_ = manager.Count()
}

func TestActivityPubActivityManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubActivityManager_ListActivityPubActivities(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	_ = manager.ListActivityPubActivities(0, 0, "", nil)
}

func TestActivityPubActivityManager_GetActivityPubActivity(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	if activityPubActivity, err := manager.GetActivityPubActivity(0, nil); err == nil {
		_ = activityPubActivity
		test.FailNow()
	}
}

func TestActivityPubActivityManager_AddActivityPubActivity(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivity, err := manager.AddActivityPubActivity(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubActivity
}

func TestActivityPubActivityManager_UpdateActivityPubActivity(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivity, err := manager.UpdateActivityPubActivity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubActivity
}

func TestActivityPubActivityManager_RemoveActivityPubActivity(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivity, err := manager.RemoveActivityPubActivity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubActivity
}

func TestActivityPubActivityManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivity := manager.Find(0)
	if activityPubActivity == nil {
		test.Fail()
	}

	_ = activityPubActivity
}

func TestActivityPubActivityManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	manager.ForEach(func(activityPubActivity IActivityPubActivity) {
		_ = activityPubActivity
	})
}

func TestActivityPubActivityManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivities := manager.Filter(func(activityPubActivity IActivityPubActivity) bool {
		return false
	})

	if activityPubActivities.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubActivities
}

func TestActivityPubActivityManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubActivityManager()

	activityPubActivities := manager.Map(func(activityPubActivity IActivityPubActivity) IActivityPubActivity {
		return activityPubActivity
	})

	if activityPubActivities.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubActivities
}
