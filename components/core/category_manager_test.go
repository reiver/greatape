package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestCategoryManager_GetName(test *testing.T) {
	manager := Conductor.CategoryManager()

	if manager.Name() != CATEGORY_MANAGER {
		test.Fail()
	}
}

func TestCategoryManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.CategoryManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestCategoryManager_Load(test *testing.T) {
	manager := Conductor.CategoryManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestCategoryManager_Reload(test *testing.T) {
	manager := Conductor.CategoryManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestCategoryManager_Count(test *testing.T) {
	manager := Conductor.CategoryManager()

	_ = manager.Count()
}

func TestCategoryManager_Exists(test *testing.T) {
	manager := Conductor.CategoryManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestCategoryManager_ListCategories(test *testing.T) {
	manager := Conductor.CategoryManager()

	_ = manager.ListCategories(0, 0, "", nil)
}

func TestCategoryManager_GetCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	if category, err := manager.GetCategory(0, nil); err == nil {
		_ = category
		test.FailNow()
	}
}

func TestCategoryManager_AddCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	category, err := manager.AddCategory(0, 0, "title", "description", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = category
}

func TestCategoryManager_UpdateCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	category, err := manager.UpdateCategory(0, 0, 0, "title", "description", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = category
}

func TestCategoryManager_RemoveCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	category, err := manager.RemoveCategory(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = category
}

func TestCategoryManager_Find(test *testing.T) {
	manager := Conductor.CategoryManager()

	category := manager.Find(0)
	if category == nil {
		test.Fail()
	}

	_ = category
}

func TestCategoryManager_ForEach(test *testing.T) {
	manager := Conductor.CategoryManager()

	manager.ForEach(func(category ICategory) {
		_ = category
	})
}

func TestCategoryManager_Filter(test *testing.T) {
	manager := Conductor.CategoryManager()

	categories := manager.Filter(func(category ICategory) bool {
		return category.Id() < 0
	})

	if categories.IsNotEmpty() {
		test.Fail()
	}

	_ = categories
}

func TestCategoryManager_Map(test *testing.T) {
	manager := Conductor.CategoryManager()

	categories := manager.Map(func(category ICategory) ICategory {
		return category
	})

	if categories.Count() != manager.Count() {
		test.Fail()
	}

	_ = categories
}

func TestCategoryManager_ListCategoriesByCategoryType(test *testing.T) {
	manager := Conductor.CategoryManager()

	_ = manager.ListCategoriesByCategoryType(0, 0, 0, "", nil)
}

func TestCategoryManager_ForEachByCategoryType(test *testing.T) {
	manager := Conductor.CategoryManager()

	manager.ForEachByCategoryType(0, func(category ICategory) {
		_ = category
	})
}

func TestCategoryManager_ListCategoriesByCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	_ = manager.ListCategoriesByCategory(0, 0, 0, "", nil)
}

func TestCategoryManager_ForEachByCategory(test *testing.T) {
	manager := Conductor.CategoryManager()

	manager.ForEachByCategory(0, func(category ICategory) {
		_ = category
	})
}
