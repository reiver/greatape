package commands

import (
	. "github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func ResetPassword(x IDispatcher, usernameOrEmail string) (IResetPasswordResult, error) {
	isEmail := REGEXP_EMAIL.MatchString(usernameOrEmail)
	if !isEmail && !UsernameIsValid(usernameOrEmail) {
		return nil, ERROR_INVALID_PARAMETERS
	}

	identities := x.FilterIdentities(func(identity IIdentity) bool {
		if isEmail {
			return identity.Email() == usernameOrEmail
		} else {
			return identity.Username() == usernameOrEmail
		}
	})

	if identities.HasExactlyOneItem() {
		identity := identities.First()
		_ = identity
	}

	return x.NewResetPasswordResult(), nil
}
