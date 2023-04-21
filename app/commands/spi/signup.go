package spi

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func Signup(x IDispatcher, username string, email string, password string) (ISignupResult, error) {
	if x.IdentityExistsWhich(func(identity IIdentity) bool {
		return identity.Username() == username || identity.Email() == email
	}) {
		return nil, ERROR_USERNAME_OR_EMAIL_ALREADY_REGISTERED
	}

	salt := x.GenerateSalt()
	hash := x.GenerateHash(password, salt)
	token := x.GenerateJwtToken()
	code := x.GenerateCode()
	privateKey, publicKey, err := x.GenerateRSAKeyPair()
	x.AssertNoError(err)

	x.Atomic(func() error {
		identity := x.AddIdentity(
			username,             // username
			EMPTY,                // phoneNumber
			false,                // phoneNumberConfirmed
			EMPTY,                // firstName
			EMPTY,                // lastName
			EMPTY,                // displayName
			email,                // email
			false,                // emailConfirmed
			EMPTY,                // avatar
			EMPTY,                // banner
			EMPTY,                // summary
			code,                 // token
			false,                // multiFactor
			hash,                 // hash
			salt,                 // salt
			publicKey,            // publicKey
			privateKey,           // privateKey
			ACL_PERMISSION_USER,  // permission
			ACL_RESTRICTION_NONE, // restriction
			NOT_SET,              // lastLogin
			0,                    // loginCount
		)

		x.AddUser(
			identity.Id(), // identityId
			EMPTY,         // github
		)

		return nil
	})

	if x.IsProductionEnvironment() {
		x.Email(email, "Confirmation Code: %s", code)
		code = "Confirmation code is sent by email."
	}

	return x.NewSignupResult(token, code), nil
}
