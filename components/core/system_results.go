package core

import . "github.com/reiver/greatape/components/contracts"

//region IDispatcher Implementation

func (dispatcher *dispatcher) NewEchoResult(document IDocument) IEchoResult {
	return NewEchoResult(document, nil)
}

func (dispatcher *dispatcher) NewCheckUsernameAvailabilityResult(isAvailable bool) ICheckUsernameAvailabilityResult {
	return NewCheckUsernameAvailabilityResult(isAvailable, nil)
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

func (dispatcher *dispatcher) NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string) IGetProfileByUserResult {
	return NewGetProfileByUserResult(username, displayName, avatar, banner, summary, github, nil)
}

func (dispatcher *dispatcher) NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string) IUpdateProfileByUserResult {
	return NewUpdateProfileByUserResult(displayName, avatar, banner, summary, github, nil)
}

func (dispatcher *dispatcher) NewLogoutResult() ILogoutResult {
	return NewLogoutResult(nil)
}

func (dispatcher *dispatcher) NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string) IWebfingerResult {
	return NewWebfingerResult(aliases, links, subject, nil)
}

func (dispatcher *dispatcher) NewGetPackagesResult(body []byte) IGetPackagesResult {
	return NewGetPackagesResult(body, nil)
}

func (dispatcher *dispatcher) NewGetActorResult(context []string, id string, followers string, following string, inbox string, outbox string, name string, preferredUsername string, type_ string, url string, icon IActivityPubMedia, image IActivityPubMedia, publicKey IActivityPubPublicKey, summary string, published string) IGetActorResult {
	return NewGetActorResult(context, id, followers, following, inbox, outbox, name, preferredUsername, type_, url, icon, image, publicKey, summary, published, nil)
}

func (dispatcher *dispatcher) NewFollowActorResult(url string) IFollowActorResult {
	return NewFollowActorResult(url, nil)
}

func (dispatcher *dispatcher) NewAuthorizeInteractionResult(uri string, success bool) IAuthorizeInteractionResult {
	return NewAuthorizeInteractionResult(uri, success, nil)
}

func (dispatcher *dispatcher) NewGetFollowersResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string) IGetFollowersResult {
	return NewGetFollowersResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (dispatcher *dispatcher) NewGetFollowingResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string) IGetFollowingResult {
	return NewGetFollowingResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (dispatcher *dispatcher) NewPostToOutboxResult(body []byte) IPostToOutboxResult {
	return NewPostToOutboxResult(body, nil)
}

func (dispatcher *dispatcher) NewGetOutboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string) IGetOutboxResult {
	return NewGetOutboxResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (dispatcher *dispatcher) NewPostToInboxResult(body []byte) IPostToInboxResult {
	return NewPostToInboxResult(body, nil)
}

func (dispatcher *dispatcher) NewGetInboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string) IGetInboxResult {
	return NewGetInboxResult(context, id, type_, totalItems, orderedItems, first, nil)
}

//endregion
