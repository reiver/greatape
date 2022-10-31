package test

import "testing"

func Test_Local_Echo(t *testing.T) {
	Run(t, apiLocal, echo)
}
