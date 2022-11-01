package core

import . "rail.town/infrastructure/components/contracts"

//region IDispatcher Implementation

func (dispatcher *dispatcher) NewEchoResult(document IDocument) IEchoResult {
	return NewEchoResult(document, nil)
}

//endregion
