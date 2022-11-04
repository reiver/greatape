package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestActivityPubOutgoingActivityManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if manager.Name() != ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER {
		test.Fail()
	}
}

func TestActivityPubOutgoingActivityManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubOutgoingActivityManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubOutgoingActivityManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubOutgoingActivityManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	_ = manager.Count()
}

func TestActivityPubOutgoingActivityManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubOutgoingActivityManager_ListActivityPubOutgoingActivities(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	_ = manager.ListActivityPubOutgoingActivities(0, 0, "", nil)
}

func TestActivityPubOutgoingActivityManager_GetActivityPubOutgoingActivity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	if activityPubOutgoingActivity, err := manager.GetActivityPubOutgoingActivity(0, nil); err == nil {
		_ = activityPubOutgoingActivity
		test.FailNow()
	}
}

func TestActivityPubOutgoingActivityManager_AddActivityPubOutgoingActivity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivity, err := manager.AddActivityPubOutgoingActivity(0, "unique_identifier", 0, "from", "to", "content", "raw", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubOutgoingActivity
}

func TestActivityPubOutgoingActivityManager_UpdateActivityPubOutgoingActivity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivity, err := manager.UpdateActivityPubOutgoingActivity(0, 0, "unique_identifier", 0, "from", "to", "content", "raw", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubOutgoingActivity
}

func TestActivityPubOutgoingActivityManager_RemoveActivityPubOutgoingActivity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivity, err := manager.RemoveActivityPubOutgoingActivity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubOutgoingActivity
}

func TestActivityPubOutgoingActivityManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivity := manager.Find(0)
	if activityPubOutgoingActivity == nil {
		test.Fail()
	}

	_ = activityPubOutgoingActivity
}

func TestActivityPubOutgoingActivityManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	manager.ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		_ = activityPubOutgoingActivity
	})
}

func TestActivityPubOutgoingActivityManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivities := manager.Filter(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) bool {
		return activityPubOutgoingActivity.Id() < 0
	})

	if activityPubOutgoingActivities.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubOutgoingActivities
}

func TestActivityPubOutgoingActivityManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	activityPubOutgoingActivities := manager.Map(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) IActivityPubOutgoingActivity {
		return activityPubOutgoingActivity
	})

	if activityPubOutgoingActivities.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubOutgoingActivities
}

func TestActivityPubOutgoingActivityManager_ListActivityPubOutgoingActivitiesByIdentity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	_ = manager.ListActivityPubOutgoingActivitiesByIdentity(0, 0, 0, "", nil)
}

func TestActivityPubOutgoingActivityManager_ForEachByIdentity(test *testing.T) {
	manager := Conductor.ActivityPubOutgoingActivityManager()

	manager.ForEachByIdentity(0, func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		_ = activityPubOutgoingActivity
	})
}
