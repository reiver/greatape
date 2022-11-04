package core

import . "rail.town/infrastructure/components/contracts"

//region IDispatcher Implementation

func (dispatcher *dispatcher) NewDocuments() IDocumentCollection {
	return NewDocuments()
}

func (dispatcher *dispatcher) NewSystemSchedules() ISystemScheduleCollection {
	return NewSystemSchedules()
}

func (dispatcher *dispatcher) NewIdentities() IIdentityCollection {
	return NewIdentities()
}

func (dispatcher *dispatcher) NewAccessControls() IAccessControlCollection {
	return NewAccessControls()
}

func (dispatcher *dispatcher) NewRemoteActivities() IRemoteActivityCollection {
	return NewRemoteActivities()
}

func (dispatcher *dispatcher) NewCategoryTypes() ICategoryTypeCollection {
	return NewCategoryTypes()
}

func (dispatcher *dispatcher) NewCategories() ICategoryCollection {
	return NewCategories()
}

func (dispatcher *dispatcher) NewUsers() IUserCollection {
	return NewUsers()
}

func (dispatcher *dispatcher) NewActivityPubObjects() IActivityPubObjectCollection {
	return NewActivityPubObjects()
}

func (dispatcher *dispatcher) NewActivityPubActivities() IActivityPubActivityCollection {
	return NewActivityPubActivities()
}

func (dispatcher *dispatcher) NewActivityPubPublicKeys() IActivityPubPublicKeyCollection {
	return NewActivityPubPublicKeys()
}

func (dispatcher *dispatcher) NewActivityPubLinks() IActivityPubLinkCollection {
	return NewActivityPubLinks()
}

func (dispatcher *dispatcher) NewActivityPubMedias() IActivityPubMediaCollection {
	return NewActivityPubMedias()
}

func (dispatcher *dispatcher) NewActivityPubIncomingActivities() IActivityPubIncomingActivityCollection {
	return NewActivityPubIncomingActivities()
}

func (dispatcher *dispatcher) NewActivityPubOutgoingActivities() IActivityPubOutgoingActivityCollection {
	return NewActivityPubOutgoingActivities()
}

func (dispatcher *dispatcher) NewSpis() ISpiCollection {
	return NewSpis()
}

//endregion
