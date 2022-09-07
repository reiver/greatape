package config

import (
	"os"
	"strconv"
)

var (
	PROTOCOL = getEnv("PROTOCOL", "http")
	DOMAIN   = getEnv("DOMAIN", "localhost")
	PORT     = getEnv("PORT", "80")

	SQLITE_DB = getEnv("SQLITE_DB", "db.sqlite")
	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "put-your-secure-jwt-secret-key-here")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP = getEnv("TOKEN_EXP", "10h")

	// Maximum allowed upload file size in megabytes.
	MAX_UPLOAD_SIZE = getEnv("MAX_UPLOAD_SIZE", "1")
	UPLOAD_PATH     = getEnv("UPLOAD_PATH", "./upload")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return fallback
}

func BodyLimit() int {
	maxFileSize, err := strconv.ParseInt(MAX_UPLOAD_SIZE, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(maxFileSize) * 1024 * 1024
}
