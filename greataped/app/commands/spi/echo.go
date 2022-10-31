package spi

import (
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

func Echo(x IDispatcher, document IDocument) (IEchoResult, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}
