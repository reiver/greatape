package validators

import (
	"regexp"

	. "github.com/reiver/greatape/components/constants"
)

func WebfingerIsValid(webfinger string) bool {
	return regexp.MustCompile(WEBFINGER).MatchString(webfinger)
}
