package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestActivityPubPublicKeyManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if manager.Name() != ACTIVITY_PUB_PUBLIC_KEY_MANAGER {
		test.Fail()
	}
}

func TestActivityPubPublicKeyManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubPublicKeyManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubPublicKeyManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubPublicKeyManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	_ = manager.Count()
}

func TestActivityPubPublicKeyManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubPublicKeyManager_ListActivityPubPublicKeys(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	_ = manager.ListActivityPubPublicKeys(0, 0, "", nil)
}

func TestActivityPubPublicKeyManager_GetActivityPubPublicKey(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	if activityPubPublicKey, err := manager.GetActivityPubPublicKey(0, nil); err == nil {
		_ = activityPubPublicKey
		test.FailNow()
	}
}

func TestActivityPubPublicKeyManager_AddActivityPubPublicKey(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKey, err := manager.AddActivityPubPublicKey(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubPublicKey
}

func TestActivityPubPublicKeyManager_UpdateActivityPubPublicKey(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKey, err := manager.UpdateActivityPubPublicKey(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubPublicKey
}

func TestActivityPubPublicKeyManager_RemoveActivityPubPublicKey(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKey, err := manager.RemoveActivityPubPublicKey(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubPublicKey
}

func TestActivityPubPublicKeyManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKey := manager.Find(0)
	if activityPubPublicKey == nil {
		test.Fail()
	}

	_ = activityPubPublicKey
}

func TestActivityPubPublicKeyManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	manager.ForEach(func(activityPubPublicKey IActivityPubPublicKey) {
		_ = activityPubPublicKey
	})
}

func TestActivityPubPublicKeyManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKeys := manager.Filter(func(activityPubPublicKey IActivityPubPublicKey) bool {
		return false
	})

	if activityPubPublicKeys.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubPublicKeys
}

func TestActivityPubPublicKeyManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubPublicKeyManager()

	activityPubPublicKeys := manager.Map(func(activityPubPublicKey IActivityPubPublicKey) IActivityPubPublicKey {
		return activityPubPublicKey
	})

	if activityPubPublicKeys.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubPublicKeys
}
