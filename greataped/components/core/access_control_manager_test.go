package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestAccessControlManager_GetName(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if manager.Name() != ACCESS_CONTROL_MANAGER {
		test.Fail()
	}
}

func TestAccessControlManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestAccessControlManager_Load(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestAccessControlManager_Reload(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestAccessControlManager_Count(test *testing.T) {
	manager := Conductor.AccessControlManager()

	_ = manager.Count()
}

func TestAccessControlManager_Exists(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestAccessControlManager_ListAccessControls(test *testing.T) {
	manager := Conductor.AccessControlManager()

	_ = manager.ListAccessControls(0, 0, "", nil)
}

func TestAccessControlManager_GetAccessControl(test *testing.T) {
	manager := Conductor.AccessControlManager()

	if accessControl, err := manager.GetAccessControl(0, nil); err == nil {
		_ = accessControl
		test.FailNow()
	}
}

func TestAccessControlManager_AddAccessControl(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControl, err := manager.AddAccessControl(0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = accessControl
}

func TestAccessControlManager_UpdateAccessControl(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControl, err := manager.UpdateAccessControl(0, 0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = accessControl
}

func TestAccessControlManager_RemoveAccessControl(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControl, err := manager.RemoveAccessControl(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = accessControl
}

func TestAccessControlManager_Find(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControl := manager.Find(0)
	if accessControl == nil {
		test.Fail()
	}

	_ = accessControl
}

func TestAccessControlManager_ForEach(test *testing.T) {
	manager := Conductor.AccessControlManager()

	manager.ForEach(func(accessControl IAccessControl) {
		_ = accessControl
	})
}

func TestAccessControlManager_Filter(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControls := manager.Filter(func(accessControl IAccessControl) bool {
		return accessControl.Id() < 0
	})

	if accessControls.IsNotEmpty() {
		test.Fail()
	}

	_ = accessControls
}

func TestAccessControlManager_Map(test *testing.T) {
	manager := Conductor.AccessControlManager()

	accessControls := manager.Map(func(accessControl IAccessControl) IAccessControl {
		return accessControl
	})

	if accessControls.Count() != manager.Count() {
		test.Fail()
	}

	_ = accessControls
}
