package validators_test

import (
	"testing"

	"github.com/reiver/greatape/app/validators"
)

func TestPasswordValidator(test *testing.T) {
	type arguments struct {
		password string
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
				password: "weak_password",
			},
		},
		{
			"Case2",
			false,
			arguments{
				password: "short",
			},
		},
		{
			"Case3",
			true,
			arguments{
				password: "G0Od@PaSsw0rD",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.PasswordIsValid(testCase.arguments.password); result != testCase.expectation {
				test.Errorf("PasswordIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
