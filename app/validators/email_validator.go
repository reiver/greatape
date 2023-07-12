package validators

import . "github.com/reiver/greatape/components/constants"

func EmailIsValid(input string) bool {
	return REGEXP_EMAIL.MatchString(input)
}
