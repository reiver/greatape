package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestActivityPubIncomingActivityManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if manager.Name() != ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER {
		test.Fail()
	}
}

func TestActivityPubIncomingActivityManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubIncomingActivityManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubIncomingActivityManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubIncomingActivityManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	_ = manager.Count()
}

func TestActivityPubIncomingActivityManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubIncomingActivityManager_ListActivityPubIncomingActivities(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	_ = manager.ListActivityPubIncomingActivities(0, 0, "", nil)
}

func TestActivityPubIncomingActivityManager_GetActivityPubIncomingActivity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	if activityPubIncomingActivity, err := manager.GetActivityPubIncomingActivity(0, nil); err == nil {
		_ = activityPubIncomingActivity
		test.FailNow()
	}
}

func TestActivityPubIncomingActivityManager_AddActivityPubIncomingActivity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivity, err := manager.AddActivityPubIncomingActivity(0, "unique_identifier", 0, "from", "to", "content", "raw", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubIncomingActivity
}

func TestActivityPubIncomingActivityManager_UpdateActivityPubIncomingActivity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivity, err := manager.UpdateActivityPubIncomingActivity(0, 0, "unique_identifier", 0, "from", "to", "content", "raw", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubIncomingActivity
}

func TestActivityPubIncomingActivityManager_RemoveActivityPubIncomingActivity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivity, err := manager.RemoveActivityPubIncomingActivity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubIncomingActivity
}

func TestActivityPubIncomingActivityManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivity := manager.Find(0)
	if activityPubIncomingActivity == nil {
		test.Fail()
	}

	_ = activityPubIncomingActivity
}

func TestActivityPubIncomingActivityManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	manager.ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		_ = activityPubIncomingActivity
	})
}

func TestActivityPubIncomingActivityManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivities := manager.Filter(func(activityPubIncomingActivity IActivityPubIncomingActivity) bool {
		return activityPubIncomingActivity.Id() < 0
	})

	if activityPubIncomingActivities.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubIncomingActivities
}

func TestActivityPubIncomingActivityManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	activityPubIncomingActivities := manager.Map(func(activityPubIncomingActivity IActivityPubIncomingActivity) IActivityPubIncomingActivity {
		return activityPubIncomingActivity
	})

	if activityPubIncomingActivities.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubIncomingActivities
}

func TestActivityPubIncomingActivityManager_ListActivityPubIncomingActivitiesByIdentity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	_ = manager.ListActivityPubIncomingActivitiesByIdentity(0, 0, 0, "", nil)
}

func TestActivityPubIncomingActivityManager_ForEachByIdentity(test *testing.T) {
	manager := Conductor.ActivityPubIncomingActivityManager()

	manager.ForEachByIdentity(0, func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		_ = activityPubIncomingActivity
	})
}
