package spi

import (
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

func Verify(x IDispatcher, email string, token string, code string) (IVerifyResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Email() == email
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_REGISTERED)
	identity := identities.First()

	if code != identity.Token() {
		return nil, ERROR_INVALID_CONFIRMATION_CODE
	}

	err := x.VerifyJwtToken(token)
	x.AssertNull(err).Or(ERROR_INVALID_TOKEN)
	token = x.GenerateJwtToken()

	x.Atomic(func() error {
		count := identity.LoginCount() + 1
		identity.UpdateLastLoginAtomic(x.Transaction(), x.UnixNano(), identity)
		identity.UpdateLoginCountAtomic(x.Transaction(), count, identity)
		identity.UpdateTokenAtomic(x.Transaction(), token, identity)
		x.IdentityManager().RefreshTokenCache(identity, token)
		return nil
	})

	return x.NewVerifyResult(token), nil
}
