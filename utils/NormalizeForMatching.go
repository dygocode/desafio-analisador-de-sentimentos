package utils

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

// Função para normalizar uma palavra para fins de comparação
func NormalizeForMatching(token string) string {
	// passar para lowercase
	tempToken := strings.ToLower(token)

	// aplicar Normalization Form Compatibility / Decomposition (NFKD)
	tokenToNFKD := norm.NFKD.String(tempToken)

	// remove pontuações no texto
	clearedToken := RemovePunctuation(tokenToNFKD)

	// remove diácriticos
	normalizedToken := RemoveDiacritics(clearedToken)

	// retorna token normalizado
	return normalizedToken
}
