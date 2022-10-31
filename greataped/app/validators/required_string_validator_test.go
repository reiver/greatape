package validators_test

import (
	"testing"

	"rail.town/infrastructure/app/validators"
)

func TestRequiredStringValidator(test *testing.T) {
	type arguments struct {
		input string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			"Case1",
			true,
			arguments{
				input: "sample",
			},
		},
		{
			"Case2",
			false,
			arguments{
				input: "",
			},
		},
		{
			"Case3",
			false,
			arguments{
				input: " ",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.RequiredStringIsValid(testCase.arguments.input); result != testCase.expectation {
				test.Errorf("RequiredStringIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
