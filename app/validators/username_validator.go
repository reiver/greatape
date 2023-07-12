package validators

import . "github.com/reiver/greatape/components/constants"

func UsernameIsValid(username string) bool {
	if !REGEXP_USERNAME.MatchString(username) {
		return false
	}

	for _, reservedUsername := range ReservedUsernames {
		if username == reservedUsername {
			return false
		}
	}

	return true
}
