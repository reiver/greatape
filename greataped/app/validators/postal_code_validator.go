package validators

import "regexp"

func PostalCodeIsValid(postalCode string) bool {
	// Optional
	if postalCode == "" {
		return true
	}

	match, err := regexp.MatchString("^\\d{10}$", postalCode)
	if err != nil || !match {
		return false
	}

	return true
}
