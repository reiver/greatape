package commands

import . "github.com/reiver/greatape/components/contracts"

func CheckUsernameAvailability(x IDispatcher, username string) (ICheckUsernameAvailabilityResult, error) {
	isAvailable := true
	if x.IdentityExistsWhich(func(identity IIdentity) bool {
		return identity.Username() == username
	}) {
		isAvailable = false
	}

	return x.NewCheckUsernameAvailabilityResult(isAvailable), nil
}
