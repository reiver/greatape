package validators_test

import (
	"testing"

	"github.com/reiver/greatape/app/validators"
)

func TestPercentValidator(test *testing.T) {
	type arguments struct {
		percent float64
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
				percent: 0,
			},
		},
		{
			"Case2",
			true,
			arguments{
				percent: 30,
			},
		},
		{
			"Case3",
			false,
			arguments{
				percent: 110,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := validators.PercentIsValid(testCase.arguments.percent); result != testCase.expectation {
				test.Errorf("PercentIsValid() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
