package utils

import "strings"

// Função para dividir o conteúdo da mensagem em palavras
func SplitContent(content string) []string {
	splitedWords := strings.Split(content, " ")
	return splitedWords
}
