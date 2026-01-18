package utils

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// RemoveDiacritics remove diácriticos de uma string
func RemoveDiacritics(token string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	normalized, _, err := transform.String(t, token)
	if err != nil {
		return token
	}

	return normalized
}
