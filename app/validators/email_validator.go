package validators

import (
	"regexp"

	. "github.com/reiver/greatape/components/constants"
)

func EmailIsValid(input string) bool {
	return regexp.MustCompile(EMAIL).MatchString(input)
}
