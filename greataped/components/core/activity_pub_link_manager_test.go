package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestActivityPubLinkManager_GetName(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if manager.Name() != ACTIVITY_PUB_LINK_MANAGER {
		test.Fail()
	}
}

func TestActivityPubLinkManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubLinkManager_Load(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestActivityPubLinkManager_Reload(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestActivityPubLinkManager_Count(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	_ = manager.Count()
}

func TestActivityPubLinkManager_Exists(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestActivityPubLinkManager_ListActivityPubLinks(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	_ = manager.ListActivityPubLinks(0, 0, "", nil)
}

func TestActivityPubLinkManager_GetActivityPubLink(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	if activityPubLink, err := manager.GetActivityPubLink(0, nil); err == nil {
		_ = activityPubLink
		test.FailNow()
	}
}

func TestActivityPubLinkManager_AddActivityPubLink(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLink, err := manager.AddActivityPubLink(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubLink
}

func TestActivityPubLinkManager_UpdateActivityPubLink(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLink, err := manager.UpdateActivityPubLink(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubLink
}

func TestActivityPubLinkManager_RemoveActivityPubLink(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLink, err := manager.RemoveActivityPubLink(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = activityPubLink
}

func TestActivityPubLinkManager_Find(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLink := manager.Find(0)
	if activityPubLink == nil {
		test.Fail()
	}

	_ = activityPubLink
}

func TestActivityPubLinkManager_ForEach(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	manager.ForEach(func(activityPubLink IActivityPubLink) {
		_ = activityPubLink
	})
}

func TestActivityPubLinkManager_Filter(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLinks := manager.Filter(func(activityPubLink IActivityPubLink) bool {
		return false
	})

	if activityPubLinks.IsNotEmpty() {
		test.Fail()
	}

	_ = activityPubLinks
}

func TestActivityPubLinkManager_Map(test *testing.T) {
	manager := Conductor.ActivityPubLinkManager()

	activityPubLinks := manager.Map(func(activityPubLink IActivityPubLink) IActivityPubLink {
		return activityPubLink
	})

	if activityPubLinks.Count() != manager.Count() {
		test.Fail()
	}

	_ = activityPubLinks
}
