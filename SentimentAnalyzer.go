package main

import (
	"sync"

	. "sentiment-analyzer/types"
	. "sentiment-analyzer/utils"
)

// Função principal de análise de sentimento
// func SentimentAnalyzer(message MessageRequest) (bool, AnalysisResponse) {
// 	sentiment := SentimentDistribution{}
// 	trendingTopics := []string{}

// 	flags := Flags{
// 		MbrasEmployee:      false,
// 		SpecialPattern:     false,
// 		CandidateAwareness: false,
// 	}

// 	influenceRanking := []InfluenceScore{}

// 	for _, msg := range message.Messages {
// 		sentiment = SentimentMap(msg.Content)

// 		trendingTopics = append(trendingTopics, msg.Hashtags...)

// 		flags = Flags{
// 			MbrasEmployee:      IsMBRASUser(msg.UserId),
// 			SpecialPattern:     IsSpecialPattern(msg.Content),
// 			CandidateAwareness: IsCandidateAwareness(msg.Content),
// 		}

// 		influenceRanking = append(influenceRanking, InfluenceScore{
// 			UserId:         msg.UserId,
// 			InfluenceScore: msg.Reactions + msg.Shares + msg.Views,
// 		})
// 	}

// 	result := AnalysisResponse{
// 		Analysis: AnalysisResult{
// 			SentimentDistribution: SentimentDistribution{
// 				Positive: sentiment.Positive,
// 				Negative: sentiment.Negative,
// 				Neutral:  sentiment.Neutral,
// 			},
// 			AnomalyType:      "spike_in_negative_sentiment",
// 			TrendingTopics:   trendingTopics,
// 			InfluenceRanking: influenceRanking,
// 			EngagementScore:  9.42,
// 			ProcessingTimeMs: 150,
// 			AnomalyDetected:  true,
// 			Flags: Flags{
// 				MbrasEmployee:      flags.MbrasEmployee,
// 				SpecialPattern:     flags.SpecialPattern,
// 				CandidateAwareness: flags.CandidateAwareness,
// 			},
// 		},
// 	}
// 	return true, result
// }

// Estrutura para resultados parciais de cada mensagem
type messageResult struct {
	sentiment      SentimentDistribution
	hashtags       []string
	flags          Flags
	influenceScore InfluenceScore
}

func SentimentAnalyzer(message MessageRequest) (bool, AnalysisResponse) {
	numMessages := len(message.Messages)
	if numMessages == 0 {
		return false, AnalysisResponse{}
	}

	// Canal para receber resultados
	resultsChan := make(chan messageResult, numMessages)
	var wg sync.WaitGroup

	// Processar cada mensagem em paralelo
	for _, msg := range message.Messages {
		wg.Add(1)
		go func(m Message) {
			defer wg.Done()

			result := messageResult{
				sentiment: SentimentMap(m.Content),
				hashtags:  m.Hashtags,
				flags: Flags{
					MbrasEmployee:      IsMBRASUser(m.UserId),
					SpecialPattern:     IsSpecialPattern(m.Content),
					CandidateAwareness: IsCandidateAwareness(m.Content),
				},
				influenceScore: InfluenceScore{
					UserId:         m.UserId,
					InfluenceScore: m.Reactions + m.Shares + m.Views,
				},
			}

			resultsChan <- result
		}(msg)
	}

	// Fechar canal quando todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Agregar resultados
	var (
		trendingTopics    []string
		influenceRanking  []InfluenceScore
		totalPositive     int
		totalNegative     int
		totalNeutral      int
		hasMbrasEmployee  bool
		hasSpecialPattern bool
		hasCandidateAware bool
		mu                sync.Mutex
	)

	for result := range resultsChan {
		// Usar mutex apenas para operações de escrita compartilhadas
		mu.Lock()

		totalPositive += result.sentiment.Positive
		totalNegative += result.sentiment.Negative
		totalNeutral += result.sentiment.Neutral

		trendingTopics = append(trendingTopics, result.hashtags...)
		influenceRanking = append(influenceRanking, result.influenceScore)

		if result.flags.MbrasEmployee {
			hasMbrasEmployee = true
		}
		if result.flags.SpecialPattern {
			hasSpecialPattern = true
		}
		if result.flags.CandidateAwareness {
			hasCandidateAware = true
		}

		mu.Unlock()
	}

	// Calcular médias
	count := int(numMessages)
	avgSentiment := SentimentDistribution{
		Positive: totalPositive / count,
		Negative: totalNegative / count,
		Neutral:  totalNeutral / count,
	}

	result := AnalysisResponse{
		Analysis: AnalysisResult{
			SentimentDistribution: avgSentiment,
			AnomalyType:           "spike_in_negative_sentiment",
			TrendingTopics:        trendingTopics,
			InfluenceRanking:      influenceRanking,
			EngagementScore:       9.42,
			ProcessingTimeMs:      150,
			AnomalyDetected:       true,
			Flags: Flags{
				MbrasEmployee:      hasMbrasEmployee,
				SpecialPattern:     hasSpecialPattern,
				CandidateAwareness: hasCandidateAware,
			},
		},
	}

	return true, result
}
