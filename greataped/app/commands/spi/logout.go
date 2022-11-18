package spi

import . "rail.town/infrastructure/components/contracts"

func Logout(x IDispatcher) (ILogoutResult, error) {
	x.Ensure(x.SignOut())
	return x.NewLogoutResult(), nil
}
