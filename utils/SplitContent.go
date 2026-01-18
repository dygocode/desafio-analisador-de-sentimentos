package utils

import "strings"

// SplitContent divide o conteúdo da mensagem em palavras
func SplitContent(content string) []string {
	splitedWords := strings.Split(content, " ")
	return splitedWords
}
