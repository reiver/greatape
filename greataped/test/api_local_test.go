package test

import "testing"

func Test_Local_Echo(t *testing.T) {
	Run(t, apiLocal, echo)
}

func Test_Local_Signup(t *testing.T) {
	Run(t, apiLocal, signup)
}

func Test_Local_Logout(t *testing.T) {
	Run(t, apiLocal, logout)
}
