package config

import (
	"os"
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
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	return fallback
}
