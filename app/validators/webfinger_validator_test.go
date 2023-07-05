package validators_test

import (
	"testing"

	"github.com/reiver/greatape/app/validators"
)

func TestWebfingerValidator(test *testing.T) {
	type arguments struct {
		webfinger string
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
				webfinger: "somebody@somewhere.xyz",
			},
		},
		{
			"Case2",
			true,
			arguments{
				webfinger: "acct:somebody@somewhere.xyz",
			},
		},
		{
			"Case3",
			true,
			arguments{
				webfinger: "acct:user@sub.domain.com",
			},
		},
		{
			"Case4",
			false,
			arguments{
				webfinger: "acct:user",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.WebfingerIsValid(testCase.arguments.webfinger); result != testCase.expectation {
				test.Errorf("WebfingerIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
