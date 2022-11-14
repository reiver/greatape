package core

import . "rail.town/infrastructure/components/contracts"

//region IDispatcher Implementation

func (dispatcher *dispatcher) NewEchoResult(document IDocument) IEchoResult {
	return NewEchoResult(document, nil)
}

func (dispatcher *dispatcher) NewSignupResult(token string, code string) ISignupResult {
	return NewSignupResult(token, code, nil)
}

func (dispatcher *dispatcher) NewVerifyResult(token string) IVerifyResult {
	return NewVerifyResult(token, nil)
}

func (dispatcher *dispatcher) NewLoginResult(username string, token string) ILoginResult {
	return NewLoginResult(username, token, nil)
}

//endregion
