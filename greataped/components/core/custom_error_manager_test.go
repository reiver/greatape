package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestCustomErrorManager_GetName(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if manager.Name() != CUSTOM_ERROR_MANAGER {
		test.Fail()
	}
}

func TestCustomErrorManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestCustomErrorManager_Load(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestCustomErrorManager_Reload(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestCustomErrorManager_Count(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	_ = manager.Count()
}

func TestCustomErrorManager_Exists(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestCustomErrorManager_ListCustomErrors(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	_ = manager.ListCustomErrors(0, 0, "", nil)
}

func TestCustomErrorManager_GetCustomError(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	if customError, err := manager.GetCustomError(0, nil); err == nil {
		_ = customError
		test.FailNow()
	}
}

func TestCustomErrorManager_AddCustomError(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customError, err := manager.AddCustomError(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = customError
}

func TestCustomErrorManager_UpdateCustomError(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customError, err := manager.UpdateCustomError(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = customError
}

func TestCustomErrorManager_RemoveCustomError(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customError, err := manager.RemoveCustomError(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = customError
}

func TestCustomErrorManager_Find(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customError := manager.Find(0)
	if customError == nil {
		test.Fail()
	}

	_ = customError
}

func TestCustomErrorManager_ForEach(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	manager.ForEach(func(customError ICustomError) {
		_ = customError
	})
}

func TestCustomErrorManager_Filter(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customErrors := manager.Filter(func(customError ICustomError) bool {
		return false
	})

	if customErrors.IsNotEmpty() {
		test.Fail()
	}

	_ = customErrors
}

func TestCustomErrorManager_Map(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	customErrors := manager.Map(func(customError ICustomError) ICustomError {
		return customError
	})

	if customErrors.Count() != manager.Count() {
		test.Fail()
	}

	_ = customErrors
}

func TestCustomErrorManager_ResolveError(test *testing.T) {
	manager := Conductor.CustomErrorManager()

	result, err := manager.ResolveError(nil, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = result
}
