package spi

import . "github.com/reiver/greatape/components/contracts"

func GetProfileByUser(x IDispatcher) (IGetProfileByUserResult, error) {
	identity := x.Identity().(IIdentity)
	user := x.GetUser(identity.Id())

	return x.NewGetProfileByUserResult(
		identity.Username(),
		identity.DisplayName(),
		identity.Avatar(),
		identity.Banner(),
		identity.Summary(),
		user.Github(),
	), nil
}
