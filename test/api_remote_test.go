package test

import "testing"

func Test_Remote_Echo(t *testing.T) {
	Run(t, apiRemote, echo)
}

func Test_Remote_Signup(t *testing.T) {
	Run(t, apiRemote, signup)
}

func Test_Remote_Logout(t *testing.T) {
	Run(t, apiRemote, logout)
}
