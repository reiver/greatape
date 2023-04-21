package constants

// noinspection GoSnakeCaseUsage, GoUnusedConst
const (
	EMPTY      = ""
	EMPTY_JSON = "{}"
	NOT_SET    = 0

	// ACL_ROLE
	ACL_PERMISSION_USER  = 1
	ACL_PERMISSION_ADMIN = 4294967295

	// ACL_RESTRICTION
	ACL_RESTRICTION_NONE          = 0
	ACL_RESTRICTION_ACCESS_DENIED = 1

	// CRON
	CRON_EVERY_10_SECONDS  = "0/10 * * * * ?"
	CRON_EVERY_MINUTE      = "0 * * * * ?"
	CRON_EVERY_TEN_MINUTES = "0 0/10 * * * ?"
	CRON_EVERY_DAY_6PM     = "0 0 18 * * ?"
	CRON_EVERY_HOUR        = "0 0 * * * ?"
)
