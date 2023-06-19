package commands

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func ResendVerificationCode(x IDispatcher, email string) (IResendVerificationCodeResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Email() == email
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_INVALID_EMAIL_FOR_RESEND_VERIFICATION_CODE)
	identity := identities.First()
	code := identity.Token()

	if len(code) >= 10 {
		return nil, ERROR_USERNAME_OR_EMAIL_ALREADY_REGISTERED
	}

	if x.IsStagingEnvironment() || x.IsProductionEnvironment() {
		x.Email(email, "resend-verification-code",
			map[string]interface{}{
				"app":  "GreatApe",
				"code": code,
			})

		code = "000000"
	}

	return x.NewResendVerificationCodeResult(code), nil
}
