package utils

import (
	. "sentiment-analyzer/types"
	"strings"
)

func SentimentMap(content string) SentimentDistribution {
	tokens := SplitContent(content)

	var result SentimentDistribution
	lexiconMatches := 0 // controla se houve match no léxico

	for _, token := range tokens {
		// ignora hashtags
		if strings.HasPrefix(token, "#") {
			continue
		}

		normalized := NormalizeForMatching(token)

		if positiveVal, ok := positive[normalized]; ok {
			result.Positive += positiveVal
			lexiconMatches++
			continue
		}

		if negativeVal, ok := negative[normalized]; ok {
			result.Negative += negativeVal
			lexiconMatches++
			continue
		}
	}

	if IsCandidateAwareness(content) {
		return SentimentDistribution{
			Positive: 0,
			Negative: 0,
			Neutral:  0,
		}
	}

	// Nenhum termo do léxico encontrado
	if lexiconMatches == 0 {
		return SentimentDistribution{
			Positive: 0,
			Negative: 0,
			Neutral:  100,
		}
	}

	return result
}

var positive = map[string]int{
	"bom":       100,
	"adorei":    100,
	"excelente": 100,
	"gostei":    100,
}

var negative = map[string]int{
	"nao":     -100,
	"ruim":    -100,
	"odiei":   -100,
	"pessimo": -100,
}

var intensifiers = map[string]int{
	"muito": 150,
	"super": 150,
}
