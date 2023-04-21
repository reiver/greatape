package spi

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func GetFollowing(x IDispatcher, username string) (IGetFollowingResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	actor := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())

	var orderedItems []string = []string{}
	return x.NewGetFollowingResult(
		ACTIVITY_STREAMS,                // context
		x.Format("%s/following", actor), // id
		ACTIVITY_PUB_ORDERED_COLLECTION, // type
		int32(len(orderedItems)),        // totalItems
		orderedItems,                    // orderedItems
		"",                              // first
	), nil
}
