package utility

import (
	"contracts"
	"fmt"
	"strings"
)

type stringUtil struct{}

func NewStringUtil() contracts.IStringUtil {
	return &stringUtil{}
}

// Contains reports whether substr is within s.
func (util *stringUtil) Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func (util *stringUtil) Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Format formats according to a format specifier and returns the resulting string.
func (util *stringUtil) Format(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
