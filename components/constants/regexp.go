package constants

import "regexp"

var (
	REGEXP_USERNAME     = regexp.MustCompile(`^[a-z0-9_\.]{5,16}$`)
	REGEXP_PASSWORD     = regexp.MustCompile(`^.{6,}$`)
	REGEXP_PHONE_NUMBER = regexp.MustCompile(`^9\d{9}$`)
	REGEXP_URL          = regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	REGEXP_EMAIL        = regexp.MustCompile(`^[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})$`)
	REGEXP_WEBFINGER    = regexp.MustCompile(`^acct:[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})$`)
)
