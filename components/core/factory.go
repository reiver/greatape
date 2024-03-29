package core

import (
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/settings"
)

type systemComponentFactory struct {
	components []ISystemComponent
}

func (factory *systemComponentFactory) Create(componentType SystemComponentType, configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ISystemComponent {
	var component ISystemComponent = nil

	switch componentType {
	case SYSTEM_COMPONENT_DOCUMENT_MANAGER:
		component = newDocumentManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_SYSTEM_SCHEDULE_MANAGER:
		component = newSystemScheduleManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_IDENTITY_MANAGER:
		component = newIdentityManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACCESS_CONTROL_MANAGER:
		component = newAccessControlManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_REMOTE_ACTIVITY_MANAGER:
		component = newRemoteActivityManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_CATEGORY_TYPE_MANAGER:
		component = newCategoryTypeManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_CATEGORY_MANAGER:
		component = newCategoryManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_USER_MANAGER:
		component = newUserManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_OBJECT_MANAGER:
		component = newActivityPubObjectManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_ACTIVITY_MANAGER:
		component = newActivityPubActivityManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_PUBLIC_KEY_MANAGER:
		component = newActivityPubPublicKeyManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_LINK_MANAGER:
		component = newActivityPubLinkManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_MEDIA_MANAGER:
		component = newActivityPubMediaManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER:
		component = newActivityPubIncomingActivityManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER:
		component = newActivityPubOutgoingActivityManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_ACTIVITY_PUB_FOLLOWER_MANAGER:
		component = newActivityPubFollowerManager(configuration, logger, dependencies...)
	case SYSTEM_COMPONENT_SPI_MANAGER:
		component = newSpiManager(configuration, logger, dependencies...)
	}

	if component != nil {
		factory.components = append(factory.components, component)
	}

	return component
}

func (factory *systemComponentFactory) Components() []ISystemComponent {
	return factory.components
}

func newSystemComponentFactory() ISystemComponentFactory {
	return &systemComponentFactory{
		components: []ISystemComponent{},
	}
}
