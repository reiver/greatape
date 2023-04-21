package spi

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func Login(x IDispatcher, email string, password string) (ILoginResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Email() == email
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_INVALID_CREDENTIALS)
	identity := identities.First()

	if len(identity.Token()) < 10 {
		return nil, ERROR_ACCOUNT_NOT_VERIFIED
	}

	if x.GenerateHash(password, identity.Salt()) != identity.Hash() {
		return nil, ERROR_INVALID_CREDENTIALS
	}

	token := x.GenerateJwtToken()

	x.Atomic(func() error {
		count := identity.LoginCount() + 1
		identity.UpdateLastLoginAtomic(x.Transaction(), x.UnixNano(), identity)
		identity.UpdateLoginCountAtomic(x.Transaction(), count, identity)
		identity.UpdateTokenAtomic(x.Transaction(), token, identity)
		x.IdentityManager().RefreshTokenCache(identity, token)
		return nil
	})

	return x.NewLoginResult(identity.Username(), token), nil
}
