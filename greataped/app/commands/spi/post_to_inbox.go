package spi

import (
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

func PostToInbox(x IDispatcher, username string) (IPostToInboxResult, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}
