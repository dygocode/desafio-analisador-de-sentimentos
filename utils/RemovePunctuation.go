package utils

import "strings"

// RemovePunctuation remove pontuação do texto
func RemovePunctuation(text string) string {
	punctuation := ".,!?;:\"()[]{}…"

	result := text

	for _, p := range punctuation {
		result = strings.ReplaceAll(result, string(p), " ")
	}

	result = strings.TrimSpace(result)
	return result
}
