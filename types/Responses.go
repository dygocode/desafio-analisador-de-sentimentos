package types

// AnalysisResponse representa a estrutura do Response
type AnalysisResponse struct {
	Analysis AnalysisResult `json:"analysis"`
}

// AnalysisResult representa o resultado da análise de sentimento
type AnalysisResult struct {
	SentimentDistribution SentimentDistribution `json:"sentiment_distribution"`
	AnomalyType           string                `json:"anomaly_type"`
	TrendingTopics        []string              `json:"trending_topics"`
	InfluenceRanking      []InfluenceScore      `json:"influence_ranking"`
	EngagementScore       float32               `json:"engagement_score"`
	ProcessingTimeMs      int                   `json:"processing_time_ms"`
	AnomalyDetected       bool                  `json:"anomaly_detected"`
	Flags                 Flags                 `json:"flags"`
}

// SentimentDistribution representa a distribuição de sentimentos
type SentimentDistribution struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

// InfluenceScore representa a pontuação de influência de um usuário
type InfluenceScore struct {
	UserID         string `json:"user_id"`
	InfluenceScore int    `json:"influence_score"`
}

// Flags representa os flags especiais associados à análise
type Flags struct {
	MbrasEmployee      bool `json:"mbras_employee"`
	SpecialPattern     bool `json:"special_pattern"`
	CandidateAwareness bool `json:"candidate_awareness"`
}
