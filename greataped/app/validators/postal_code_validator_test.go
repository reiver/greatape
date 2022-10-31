package validators_test

import (
	"testing"

	"rail.town/infrastructure/app/validators"
)

func TestPostalCodeValidator(test *testing.T) {
	type arguments struct {
		postalCode string
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
				postalCode: "",
			},
		},
		{
			"Case2",
			false,
			arguments{
				postalCode: "74ah8e",
			},
		},
		{
			"Case3",
			true,
			arguments{
				postalCode: "0374829252",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.PostalCodeIsValid(testCase.arguments.postalCode); result != testCase.expectation {
				test.Errorf("PostalCodeIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
