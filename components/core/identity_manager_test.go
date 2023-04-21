package core_test

import (
	"testing"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
)

func TestIdentityManager_GetName(test *testing.T) {
	manager := Conductor.IdentityManager()

	if manager.Name() != IDENTITY_MANAGER {
		test.Fail()
	}
}

func TestIdentityManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.IdentityManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestIdentityManager_Load(test *testing.T) {
	manager := Conductor.IdentityManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestIdentityManager_Reload(test *testing.T) {
	manager := Conductor.IdentityManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestIdentityManager_Count(test *testing.T) {
	manager := Conductor.IdentityManager()

	_ = manager.Count()
}

func TestIdentityManager_Exists(test *testing.T) {
	manager := Conductor.IdentityManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestIdentityManager_ListIdentities(test *testing.T) {
	manager := Conductor.IdentityManager()

	_ = manager.ListIdentities(0, 0, "", nil)
}

func TestIdentityManager_GetIdentity(test *testing.T) {
	manager := Conductor.IdentityManager()

	if identity, err := manager.GetIdentity(0, nil); err == nil {
		_ = identity
		test.FailNow()
	}
}

func TestIdentityManager_AddIdentity(test *testing.T) {
	manager := Conductor.IdentityManager()

	identity, err := manager.AddIdentity("username", "phone_number", true, "first_name", "last_name", "display_name", "email", true, "avatar", "banner", "summary", "token", true, "hash", "salt", "public_key", "private_key", 0, 0, 0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = identity
}

func TestIdentityManager_UpdateIdentity(test *testing.T) {
	manager := Conductor.IdentityManager()

	identity, err := manager.UpdateIdentity(0, "username", "phone_number", true, "first_name", "last_name", "display_name", "email", true, "avatar", "banner", "summary", "token", true, "hash", "salt", "public_key", "private_key", 0, 0, 0, 0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = identity
}

func TestIdentityManager_RemoveIdentity(test *testing.T) {
	manager := Conductor.IdentityManager()

	identity, err := manager.RemoveIdentity(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = identity
}

func TestIdentityManager_Find(test *testing.T) {
	manager := Conductor.IdentityManager()

	identity := manager.Find(0)
	if identity == nil {
		test.Fail()
	}

	_ = identity
}

func TestIdentityManager_ForEach(test *testing.T) {
	manager := Conductor.IdentityManager()

	manager.ForEach(func(identity IIdentity) {
		_ = identity
	})
}

func TestIdentityManager_Filter(test *testing.T) {
	manager := Conductor.IdentityManager()

	identities := manager.Filter(func(identity IIdentity) bool {
		return identity.Id() < 0
	})

	if identities.IsNotEmpty() {
		test.Fail()
	}

	_ = identities
}

func TestIdentityManager_Map(test *testing.T) {
	manager := Conductor.IdentityManager()

	identities := manager.Map(func(identity IIdentity) IIdentity {
		return identity
	})

	if identities.Count() != manager.Count() {
		test.Fail()
	}

	_ = identities
}
