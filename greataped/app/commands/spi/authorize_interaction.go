package spi

import . "rail.town/infrastructure/components/contracts"

func AuthorizeInteraction(x IDispatcher, uri string) (IAuthorizeInteractionResult, error) {
	return x.NewAuthorizeInteractionResult(
		uri,  // uri
		true, // success
	), nil
}
