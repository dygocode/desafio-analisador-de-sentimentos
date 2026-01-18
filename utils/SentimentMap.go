package utils

import (
	types "sentiment-analyzer/types"
	"strings"
)

// SentimentMap mapeia o sentimento-valor do conteúdo da mensagem
func SentimentMap(content string) types.SentimentDistribution {
	tokens := SplitContent(content)

	var result types.SentimentDistribution

	lexiconMatches := 0

	for _, token := range tokens {
		// ignora hashtags
		if strings.HasPrefix(token, "#") {
			continue
		}

		normalized := NormalizeForMatching(token)

		// Verifica termos positivos
		if positiveVal, ok := positive[normalized]; ok {
			result.Positive += positiveVal
			lexiconMatches++
			continue
		}

		// Verifica termos negativos (usa valor absoluto)
		if negativeVal, ok := negative[normalized]; ok {
			result.Negative += negativeVal // Valor absoluto!
			lexiconMatches++
			continue
		}
	}

	// Retorna neutro se for awareness
	if IsCandidateAwareness(content) {
		return types.SentimentDistribution{
			Positive: 0,
			Negative: 0,
			Neutral:  0,
		}
	}

	// Nenhum termo do léxico encontrado = neutro
	if lexiconMatches == 0 {
		return types.SentimentDistribution{
			Positive: 0,
			Negative: 0,
			Neutral:  100,
		}
	}
	positiveScore := filterLimit(result.Positive, 0, 100)
	negativeScore := filterLimit(result.Negative, -100, 0)

	return types.SentimentDistribution{
		Positive: positiveScore,
		Negative: negativeScore,
		Neutral:  0,
	}
}

func filterLimit(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

var positive = map[string]int{
	"bom":       100,
	"adorei":    100,
	"excelente": 100,
	"gostei":    100,
	"feliz":     100,
	"otimo":     100,
	"maravilha": 100,
	"perfeito":  100,
}

var negative = map[string]int{
	"nao":      -100,
	"ruim":     -100,
	"odiei":    -100,
	"pessimo":  -100,
	"terrivel": -100,
	"horrivel": -100,
	"triste":   -100,
}

var intensifiers = map[string]int{
	"muito":  150,
	"super":  150,
	"demais": 150,
}
