package spi

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToInbox(x IDispatcher, username string) (IPostToInboxResult, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}
