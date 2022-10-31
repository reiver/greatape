package constants

import . "github.com/xeronith/diamante/contracts/localization"

// noinspection GoUnusedGlobalVariable
var Errors = Resource{
	// SYSTEM_ERRORS
	ERROR_MESSAGE_INITIALIZE:                 "initialize",
	ERROR_MESSAGE_NOT_IMPLEMENTED:            "not_implemented",
	ERROR_MESSAGE_OPERATION_FAILED:           "operation_failed",
	ERROR_MESSAGE_OPERATION_NOT_SUPPORTED:    "operation_not_supported",
	ERROR_MESSAGE_UNRESOLVED_DEPENDENCIES:    "unresolved_dependencies",
	ERROR_MESSAGE_SYSTEM_COMPONENT_NOT_FOUND: "system_component_not_found",
	ERROR_MESSAGE_DOCUMENT_NOT_FOUND:         "document_not_found",
	ERROR_MESSAGE_SYSTEM_SCHEDULE_NOT_FOUND:  "system_schedule_not_found",
	ERROR_MESSAGE_IDENTITY_NOT_FOUND:         "identity_not_found",
	ERROR_MESSAGE_ACCESS_CONTROL_NOT_FOUND:   "access_control_not_found",
	ERROR_MESSAGE_REMOTE_ACTIVITY_NOT_FOUND:  "remote_activity_not_found",
	ERROR_MESSAGE_CATEGORY_TYPE_NOT_FOUND:    "category_type_not_found",
	ERROR_MESSAGE_CATEGORY_NOT_FOUND:         "category_not_found",
	ERROR_MESSAGE_USER_NOT_FOUND:             "user_not_found",
	ERROR_MESSAGE_SPI_NOT_FOUND:              "spi_not_found",
	ERROR_MESSAGE_CUSTOM_ERROR_NOT_FOUND:     "custom_error_not_found",
	ERROR_MESSAGE_UNKNOWN_DOCUMENT:           "unknown_document",
	ERROR_MESSAGE_UNKNOWN_SYSTEM_SCHEDULE:    "unknown_system_schedule",
	ERROR_MESSAGE_UNKNOWN_IDENTITY:           "unknown_identity",
	ERROR_MESSAGE_UNKNOWN_ACCESS_CONTROL:     "unknown_access_control",
	ERROR_MESSAGE_UNKNOWN_REMOTE_ACTIVITY:    "unknown_remote_activity",
	ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE:      "unknown_category_type",
	ERROR_MESSAGE_UNKNOWN_CATEGORY:           "unknown_category",
	ERROR_MESSAGE_UNKNOWN_USER:               "unknown_user",
	ERROR_MESSAGE_UNKNOWN_SPI:                "unknown_spi",
	ERROR_MESSAGE_UNKNOWN_CUSTOM_ERROR:       "unknown_custom_error",
	ERROR_MESSAGE_INVALID_ID:                 "invalid_id",
	ERROR_MESSAGE_INVALID_PARAMETERS:         "invalid_parameters",
	// CUSTOM_ERRORS
	ERROR_MESSAGE_DATA_INTEGRITY_VIOLATION:             "data_integrity_violation",
	ERROR_MESSAGE_INVALID_STATE:                        "invalid_state",
	ERROR_MESSAGE_USER_NOT_REGISTERED:                  "user_not_registered",
	ERROR_MESSAGE_USERNAME_OR_EMAIL_ALREADY_REGISTERED: "username_or_email_already_registered",
	ERROR_MESSAGE_ACCOUNT_NOT_VERIFIED:                 "account_not_verified",
	ERROR_MESSAGE_ACCOUNT_BLOCKED:                      "account_blocked",
	ERROR_MESSAGE_INVALID_TOKEN:                        "invalid_token",
	ERROR_MESSAGE_INVALID_CONFIRMATION_CODE:            "invalid_confirmation_code",
	ERROR_MESSAGE_PERMISSION_DENIED:                    "permission_denied",
	ERROR_MESSAGE_INVALID_PERSON_KIND:                  "invalid_person_kind",
	ERROR_MESSAGE_INVALID_CREDENTIALS:                  "invalid_credentials",
}

func init() {
	// CUSTOM_ERRORS
	Errors[ERROR_MESSAGE_DATA_INTEGRITY_VIOLATION] = "data_integrity_violation"
}