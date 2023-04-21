package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestActivityPubMediaManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if manager.Name() != ACTIVITY_PUB_MEDIA_MANAGER {
		test.Fail()
	}
}

func TestActivityPubMediaManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubMediaManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubMediaManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubMediaManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	_ = manager.Count()
}

func TestActivityPubMediaManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubMediaManager_ListActivityPubMedias(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	_ = manager.ListActivityPubMedias(0, 0, "", nil)
}

func TestActivityPubMediaManager_GetActivityPubMedia(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	if activityPubMedia, err := manager.GetActivityPubMedia(0, nil); err == nil {
		_ = activityPubMedia
		test.FailNow()
	}
}

func TestActivityPubMediaManager_AddActivityPubMedia(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedia, err := manager.AddActivityPubMedia(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubMedia
}

func TestActivityPubMediaManager_UpdateActivityPubMedia(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedia, err := manager.UpdateActivityPubMedia(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubMedia
}

func TestActivityPubMediaManager_RemoveActivityPubMedia(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedia, err := manager.RemoveActivityPubMedia(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubMedia
}

func TestActivityPubMediaManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedia := manager.Find(0)
	if activityPubMedia == nil {
		test.Fail()
	}

	_ = activityPubMedia
}

func TestActivityPubMediaManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	manager.ForEach(func(activityPubMedia IActivityPubMedia) {
		_ = activityPubMedia
	})
}

func TestActivityPubMediaManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedias := manager.Filter(func(activityPubMedia IActivityPubMedia) bool {
		return false
	})

	if activityPubMedias.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubMedias
}

func TestActivityPubMediaManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubMediaManager()

	activityPubMedias := manager.Map(func(activityPubMedia IActivityPubMedia) IActivityPubMedia {
		return activityPubMedia
	})

	if activityPubMedias.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubMedias
}
