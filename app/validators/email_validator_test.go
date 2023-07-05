package validators_test

import (
	"testing"

	"github.com/reiver/greatape/app/validators"
)

func TestEmailValidator(test *testing.T) {
	type arguments struct {
		email string
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
				email: "",
			},
		},
		{
			"Case2",
			false,
			arguments{
				email: "user",
			},
		},
		{
			"Case3",
			true,
			arguments{
				email: "user@domain.com",
			},
		},
		{
			"Case4",
			true,
			arguments{
				email: "user+plus@gmail.com",
			},
		},
		{
			"Case5",
			true,
			arguments{
				email: "susan235@gmail.com",
			},
		},
		{
			"Case6",
			true,
			arguments{
				email: "new_user@icloud.com",
			},
		},
		{
			"Case7",
			true,
			arguments{
				email: "someone@somewhere.co.uk",
			},
		},
		{
			"Case8",
			true,
			arguments{
				email: "north.star@space.social",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.EmailIsValid(testCase.arguments.email); result != testCase.expectation {
				test.Errorf("EmailIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
