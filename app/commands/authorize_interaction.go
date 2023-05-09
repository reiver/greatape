package commands

import . "github.com/reiver/greatape/components/contracts"

func AuthorizeInteraction(x IDispatcher, uri string) (IAuthorizeInteractionResult, error) {
	return x.NewAuthorizeInteractionResult(
		uri,  // uri
		true, // success
	), nil
}
