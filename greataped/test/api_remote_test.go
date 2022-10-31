package test

import "testing"

func Test_Remote_Echo(t *testing.T) {
	Run(t, apiRemote, echo)
}
