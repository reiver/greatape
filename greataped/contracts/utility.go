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
	Split(s, sep string) []string
}
