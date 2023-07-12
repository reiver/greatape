package validators

import . "github.com/reiver/greatape/components/constants"

func WebfingerIsValid(webfinger string) bool {
	return REGEXP_WEBFINGER.MatchString(webfinger)
}
