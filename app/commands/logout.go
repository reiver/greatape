package commands

import . "github.com/reiver/greatape/components/contracts"

func Logout(x IDispatcher) (ILogoutResult, error) {
	x.Ensure(x.SignOut())
	return x.NewLogoutResult(), nil
}
