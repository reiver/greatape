package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestUserManager_GetName(test *testing.T) {
	manager := Conductor.UserManager()

	if manager.Name() != USER_MANAGER {
		test.Fail()
	}
}

func TestUserManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.UserManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestUserManager_Load(test *testing.T) {
	manager := Conductor.UserManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestUserManager_Reload(test *testing.T) {
	manager := Conductor.UserManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestUserManager_Count(test *testing.T) {
	manager := Conductor.UserManager()

	_ = manager.Count()
}

func TestUserManager_Exists(test *testing.T) {
	manager := Conductor.UserManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestUserManager_ListUsers(test *testing.T) {
	manager := Conductor.UserManager()

	_ = manager.ListUsers(0, 0, "", nil)
}

func TestUserManager_GetUser(test *testing.T) {
	manager := Conductor.UserManager()

	if user, err := manager.GetUser(0, nil); err == nil {
		_ = user
		test.FailNow()
	}
}

func TestUserManager_AddUser(test *testing.T) {
	manager := Conductor.UserManager()

	user, err := manager.AddUser(0, "github", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = user
}

func TestUserManager_UpdateUser(test *testing.T) {
	manager := Conductor.UserManager()

	user, err := manager.UpdateUser(0, "github", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = user
}

func TestUserManager_RemoveUser(test *testing.T) {
	manager := Conductor.UserManager()

	user, err := manager.RemoveUser(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = user
}

func TestUserManager_Find(test *testing.T) {
	manager := Conductor.UserManager()

	user := manager.Find(0)
	if user == nil {
		test.Fail()
	}

	_ = user
}

func TestUserManager_ForEach(test *testing.T) {
	manager := Conductor.UserManager()

	manager.ForEach(func(user IUser) {
		_ = user
	})
}

func TestUserManager_Filter(test *testing.T) {
	manager := Conductor.UserManager()

	users := manager.Filter(func(user IUser) bool {
		return user.Id() < 0
	})

	if users.IsNotEmpty() {
		test.Fail()
	}

	_ = users
}

func TestUserManager_Map(test *testing.T) {
	manager := Conductor.UserManager()

	users := manager.Map(func(user IUser) IUser {
		return user
	})

	if users.Count() != manager.Count() {
		test.Fail()
	}

	_ = users
}
