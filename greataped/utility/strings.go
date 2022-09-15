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

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func (util *stringUtil) Split(s, sep string) []string {
	return strings.Split(s, sep)
}
