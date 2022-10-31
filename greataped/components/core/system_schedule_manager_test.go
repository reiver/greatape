package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestSystemScheduleManager_GetName(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if manager.Name() != SYSTEM_SCHEDULE_MANAGER {
		test.Fail()
	}
}

func TestSystemScheduleManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestSystemScheduleManager_Load(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestSystemScheduleManager_Reload(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestSystemScheduleManager_Count(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	_ = manager.Count()
}

func TestSystemScheduleManager_Exists(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestSystemScheduleManager_ListSystemSchedules(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	_ = manager.ListSystemSchedules(0, 0, "", nil)
}

func TestSystemScheduleManager_GetSystemSchedule(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	if systemSchedule, err := manager.GetSystemSchedule(0, nil); err == nil {
		_ = systemSchedule
		test.FailNow()
	}
}

func TestSystemScheduleManager_AddSystemSchedule(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedule, err := manager.AddSystemSchedule(true, "config", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = systemSchedule
}

func TestSystemScheduleManager_UpdateSystemSchedule(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedule, err := manager.UpdateSystemSchedule(0, true, "config", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = systemSchedule
}

func TestSystemScheduleManager_RemoveSystemSchedule(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedule, err := manager.RemoveSystemSchedule(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = systemSchedule
}

func TestSystemScheduleManager_Find(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedule := manager.Find(0)
	if systemSchedule == nil {
		test.Fail()
	}

	_ = systemSchedule
}

func TestSystemScheduleManager_ForEach(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	manager.ForEach(func(systemSchedule ISystemSchedule) {
		_ = systemSchedule
	})
}

func TestSystemScheduleManager_Filter(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedules := manager.Filter(func(systemSchedule ISystemSchedule) bool {
		return systemSchedule.Id() < 0
	})

	if systemSchedules.IsNotEmpty() {
		test.Fail()
	}

	_ = systemSchedules
}

func TestSystemScheduleManager_Map(test *testing.T) {
	manager := Conductor.SystemScheduleManager()

	systemSchedules := manager.Map(func(systemSchedule ISystemSchedule) ISystemSchedule {
		return systemSchedule
	})

	if systemSchedules.Count() != manager.Count() {
		test.Fail()
	}

	_ = systemSchedules
}
