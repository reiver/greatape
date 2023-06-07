package sms

import "fmt"

func handler(receiver, message string, _ map[string]interface{}) error {
	return fmt.Errorf("not_implemented %s %s", receiver, message)
}
