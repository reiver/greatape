package contracts

type IStringUtil interface {
	// Contains reports whether substr is within s.
	Contains(string, string) bool
	// Replace returns a copy of the string s with the first n
	// non-overlapping instances of old replaced by new.
	// If old is empty, it matches at the beginning of the string
	// and after each UTF-8 sequence, yielding up to k+1 replacements
	// for a k-rune string.
	// If n < 0, there is no limit on the number of replacements.
	Replace(s, old, new string, n int) string
	// Format formats according to a format specifier and returns the resulting string.
	Format(format string, a ...any) string
}
