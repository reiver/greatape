package spi

import (
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

func GetInbox(x IDispatcher, username string) (IGetInboxResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	actor := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())

	var orderedItems ActivityPubActivities
	return x.NewGetInboxResult(
		ACTIVITY_STREAMS,                // context
		x.Format("%s/inbox", actor),     // id
		ACTIVITY_PUB_ORDERED_COLLECTION, // type
		0,                               // totalItems
		orderedItems,                    // orderedItems
		"",                              // first
	), nil
}
