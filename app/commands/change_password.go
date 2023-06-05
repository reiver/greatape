package commands

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func ChangePassword(x IDispatcher, currentPassword string, newPassword string) (IChangePasswordResult, error) {
	identity := x.Identity().(IIdentity)
	if len(identity.Token()) < 10 {
		return nil, ERROR_ACCOUNT_NOT_VERIFIED
	}

	if x.GenerateHash(currentPassword, identity.Salt()) != identity.Hash() {
		return nil, ERROR_INVALID_CURRENT_PASSWORD_FOR_CHANGE_PASSWORD
	}

	hash := x.GenerateHash(newPassword, identity.Salt())
	token := x.GenerateJwtToken()

	x.Atomic(func() error {
		identity.UpdateHashAtomic(x.Transaction(), hash, identity)
		identity.UpdateTokenAtomic(x.Transaction(), token, identity)
		x.IdentityManager().RefreshTokenCache(identity, token)
		return nil
	})

	return x.NewChangePasswordResult(), nil
}
