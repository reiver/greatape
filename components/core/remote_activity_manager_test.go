package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestRemoteActivityManager_GetName(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if manager.Name() != REMOTE_ACTIVITY_MANAGER {
		test.Fail()
	}
}

func TestRemoteActivityManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestRemoteActivityManager_Load(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestRemoteActivityManager_Reload(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestRemoteActivityManager_Count(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	_ = manager.Count()
}

func TestRemoteActivityManager_Exists(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestRemoteActivityManager_ListRemoteActivities(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	_ = manager.ListRemoteActivities(0, 0, "", nil)
}

func TestRemoteActivityManager_GetRemoteActivity(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	if remoteActivity, err := manager.GetRemoteActivity(0, nil); err == nil {
		_ = remoteActivity
		test.FailNow()
	}
}

func TestRemoteActivityManager_AddRemoteActivity(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivity, err := manager.AddRemoteActivity("entry_point", 0, true, "error_message", "remote_address", "user_agent", 0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = remoteActivity
}

func TestRemoteActivityManager_UpdateRemoteActivity(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivity, err := manager.UpdateRemoteActivity(0, "entry_point", 0, true, "error_message", "remote_address", "user_agent", 0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = remoteActivity
}

func TestRemoteActivityManager_RemoveRemoteActivity(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivity, err := manager.RemoveRemoteActivity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = remoteActivity
}

func TestRemoteActivityManager_Find(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivity := manager.Find(0)
	if remoteActivity == nil {
		test.Fail()
	}

	_ = remoteActivity
}

func TestRemoteActivityManager_ForEach(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	manager.ForEach(func(remoteActivity IRemoteActivity) {
		_ = remoteActivity
	})
}

func TestRemoteActivityManager_Filter(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivities := manager.Filter(func(remoteActivity IRemoteActivity) bool {
		return remoteActivity.Id() < 0
	})

	if remoteActivities.IsNotEmpty() {
		test.Fail()
	}

	_ = remoteActivities
}

func TestRemoteActivityManager_Map(test *testing.T) {
	manager := Conductor.RemoteActivityManager()

	remoteActivities := manager.Map(func(remoteActivity IRemoteActivity) IRemoteActivity {
		return remoteActivity
	})

	if remoteActivities.Count() != manager.Count() {
		test.Fail()
	}

	_ = remoteActivities
}
