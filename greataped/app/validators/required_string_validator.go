package validators

import "strings"

func RequiredStringIsValid(input string) bool {
	return strings.TrimSpace(input) == ""
}
