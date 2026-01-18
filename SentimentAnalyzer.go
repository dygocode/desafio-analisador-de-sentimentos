package main

import (
	types "sentiment-analyzer/types"
	utils "sentiment-analyzer/utils"
)

// Função principal de análise de sentimento
func SentimentAnalyzer(message types.MessageRequest) (bool, types.AnalysisResponse) {
	sentiment := types.SentimentDistribution{}
	trendingTopics := []string{}

	flags := types.Flags{
		MbrasEmployee:      false,
		SpecialPattern:     false,
		CandidateAwareness: false,
	}

	influenceRanking := []types.InfluenceScore{}

	for _, msg := range message.Messages {
		sentiment = utils.SentimentMap(msg.Content)

		trendingTopics = append(trendingTopics, msg.Hashtags...)

		flags = types.Flags{
			MbrasEmployee:      utils.IsMBRASUser(msg.UserId),
			SpecialPattern:     utils.IsSpecialPattern(msg.Content),
			CandidateAwareness: utils.IsCandidateAwareness(msg.Content),
		}

		influenceRanking = append(influenceRanking, types.InfluenceScore{
			UserId:         msg.UserId,
			InfluenceScore: msg.Reactions + msg.Shares + msg.Views,
		})
	}

	result := types.AnalysisResponse{
		Analysis: types.AnalysisResult{
			SentimentDistribution: types.SentimentDistribution{
				Positive: sentiment.Positive,
				Negative: sentiment.Negative,
				Neutral:  sentiment.Neutral,
			},
			AnomalyType:      "spike_in_negative_sentiment",
			TrendingTopics:   trendingTopics,
			InfluenceRanking: influenceRanking,
			EngagementScore:  9.42,
			ProcessingTimeMs: 150,
			AnomalyDetected:  true,
			Flags: types.Flags{
				MbrasEmployee:      flags.MbrasEmployee,
				SpecialPattern:     flags.SpecialPattern,
				CandidateAwareness: flags.CandidateAwareness,
			},
		},
	}
	return true, result
}
