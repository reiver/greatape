package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestActivityPubObjectManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if manager.Name() != ACTIVITY_PUB_OBJECT_MANAGER {
		test.Fail()
	}
}

func TestActivityPubObjectManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubObjectManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubObjectManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubObjectManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	_ = manager.Count()
}

func TestActivityPubObjectManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubObjectManager_ListActivityPubObjects(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	_ = manager.ListActivityPubObjects(0, 0, "", nil)
}

func TestActivityPubObjectManager_GetActivityPubObject(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	if activityPubObject, err := manager.GetActivityPubObject(0, nil); err == nil {
		_ = activityPubObject
		test.FailNow()
	}
}

func TestActivityPubObjectManager_AddActivityPubObject(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObject, err := manager.AddActivityPubObject(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubObject
}

func TestActivityPubObjectManager_UpdateActivityPubObject(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObject, err := manager.UpdateActivityPubObject(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubObject
}

func TestActivityPubObjectManager_RemoveActivityPubObject(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObject, err := manager.RemoveActivityPubObject(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubObject
}

func TestActivityPubObjectManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObject := manager.Find(0)
	if activityPubObject == nil {
		test.Fail()
	}

	_ = activityPubObject
}

func TestActivityPubObjectManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	manager.ForEach(func(activityPubObject IActivityPubObject) {
		_ = activityPubObject
	})
}

func TestActivityPubObjectManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObjects := manager.Filter(func(activityPubObject IActivityPubObject) bool {
		return false
	})

	if activityPubObjects.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubObjects
}

func TestActivityPubObjectManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubObjectManager()

	activityPubObjects := manager.Map(func(activityPubObject IActivityPubObject) IActivityPubObject {
		return activityPubObject
	})

	if activityPubObjects.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubObjects
}
