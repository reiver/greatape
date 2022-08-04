package config

import (
	"os"
)

var (
	PROTOCOL = getEnv("PROTOCOL", "http")
	DOMAIN   = getEnv("DOMAIN", "localhost")
	PORT     = getEnv("PORT", "80")

	EXTERNAL_PROTOCOL = getEnv("EXTERNAL_PROTOCOL", "http")
	EXTERNAL_DOMAIN   = getEnv("EXTERNAL_DOMAIN", "localhost")

	SQLITE_DB = getEnv("SQLITE_DB", "db.sqlite")
	// TOKENKEY returns the jwt token secret
	TOKENKEY = getEnv("TOKEN_KEY", "laksdjflkasjfwj92jfslj2qu0-9apsoifjk")
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
