package validators

import (
	"regexp"

	. "github.com/reiver/greatape/components/constants"
)

func UsernameIsValid(username string) bool {
	if !regexp.MustCompile(USERNAME).MatchString(username) {
		return false
	}

	for _, reservedUsername := range ReservedUsernames {
		if username == reservedUsername {
			return false
		}
	}

	return true
}
