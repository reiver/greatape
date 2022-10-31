package sms

import "fmt"

func handler(receiver, message string) error {
	return fmt.Errorf("not_implemented %s %s", receiver, message)
}
