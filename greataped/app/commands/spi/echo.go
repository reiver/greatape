package spi

import . "rail.town/infrastructure/components/contracts"

func Echo(x IDispatcher, document IDocument) (IEchoResult, error) {
	return x.NewEchoResult(document), nil
}
