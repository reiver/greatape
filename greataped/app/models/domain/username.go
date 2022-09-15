package domain

import (
	"fmt"
	"strings"
)

type Username string

func (u Username) IsFederated() bool {
	return strings.Contains(string(u), "@")
}

func (u Username) IsEmpty() bool {
	return strings.TrimSpace(string(u)) == ""
}

func (u Username) Webfinger() string {
	username := string(u)
	parts := strings.Split(username, "@")
	return fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s", parts[1], username)
}

func (u Username) String() string {
	return string(u)
}
