package strings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToCamel converts a string to camel case
func ToCamel(s string) string {
	camel := toCamel(s)
	r, w := utf8.DecodeRuneInString(camel)
	if !unicode.IsUpper(r) {
		r = unicode.ToUpper(r)
	}
	return string(r) + camel[w:]
}

// ToLowerCamel converts a string to camel case
// where first word is always lowercase
func ToLowerCamel(s string) string {
	camel := toCamel(s)
	r, w := utf8.DecodeRuneInString(camel)
	if !unicode.IsLower(r) {
		r = unicode.ToLower(r)
	}

	return string(r) + camel[w:]
}

func toCamel(s string) string {
	var (
		sb   = &strings.Builder{}
		last rune
		wc   int
	)

	for len(s) > 0 {
		r, w := utf8.DecodeRuneInString(s)
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			// last is lower and current is upper
			// e.g. Htt[pS]erver
			if unicode.IsLower(last) && unicode.IsUpper(r) {
				wc++
			}

			if !unicode.IsLetter(last) &&
				!unicode.IsNumber(r) {
				r = unicode.ToUpper(r)
				wc++
			}

			sb.WriteRune(r)
		}

		s, last = s[w:], r
	}

	// convert to lowercase when only one word
	if wc == 1 {
		return strings.ToLower(sb.String())
	}

	return sb.String()
}
