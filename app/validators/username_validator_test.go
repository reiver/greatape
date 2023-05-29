package validators_test

import (
	"testing"

	"github.com/reiver/greatape/app/validators"
)

func TestUsernameValidator(test *testing.T) {
	type arguments struct {
		username string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			"Case1",
			false,
			arguments{
				username: "",
			},
		},
		{
			"Case2",
			false,
			arguments{
				username: "admin",
			},
		},
		{
			"Case3",
			true,
			arguments{
				username: "johnny",
			},
		},
		{
			"Case4",
			false,
			arguments{
				username: "webmaster",
			},
		},
		{
			"Case5",
			true,
			arguments{
				username: "susan235",
			},
		},
		{
			"Case6",
			true,
			arguments{
				username: "new_user",
			},
		},
		{
			"Case7",
			true,
			arguments{
				username: "someone",
			},
		},
		{
			"Case8",
			true,
			arguments{
				username: "north10star",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.UsernameIsValid(testCase.arguments.username); result != testCase.expectation {
				test.Errorf("UsernameIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
