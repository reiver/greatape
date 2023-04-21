package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestCategoryTypeManager_GetName(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if manager.Name() != CATEGORY_TYPE_MANAGER {
		test.Fail()
	}
}

func TestCategoryTypeManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestCategoryTypeManager_Load(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestCategoryTypeManager_Reload(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestCategoryTypeManager_Count(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	_ = manager.Count()
}

func TestCategoryTypeManager_Exists(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestCategoryTypeManager_ListCategoryTypes(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	_ = manager.ListCategoryTypes(0, 0, "", nil)
}

func TestCategoryTypeManager_GetCategoryType(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	if categoryType, err := manager.GetCategoryType(0, nil); err == nil {
		_ = categoryType
		test.FailNow()
	}
}

func TestCategoryTypeManager_AddCategoryType(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryType, err := manager.AddCategoryType("description", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = categoryType
}

func TestCategoryTypeManager_UpdateCategoryType(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryType, err := manager.UpdateCategoryType(0, "description", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = categoryType
}

func TestCategoryTypeManager_RemoveCategoryType(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryType, err := manager.RemoveCategoryType(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = categoryType
}

func TestCategoryTypeManager_Find(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryType := manager.Find(0)
	if categoryType == nil {
		test.Fail()
	}

	_ = categoryType
}

func TestCategoryTypeManager_ForEach(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	manager.ForEach(func(categoryType ICategoryType) {
		_ = categoryType
	})
}

func TestCategoryTypeManager_Filter(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryTypes := manager.Filter(func(categoryType ICategoryType) bool {
		return categoryType.Id() < 0
	})

	if categoryTypes.IsNotEmpty() {
		test.Fail()
	}

	_ = categoryTypes
}

func TestCategoryTypeManager_Map(test *testing.T) {
	manager := Conductor.CategoryTypeManager()

	categoryTypes := manager.Map(func(categoryType ICategoryType) ICategoryType {
		return categoryType
	})

	if categoryTypes.Count() != manager.Count() {
		test.Fail()
	}

	_ = categoryTypes
}
