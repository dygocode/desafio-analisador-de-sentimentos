package types

// Estruturas do Response
type AnalysisResponse struct {
	Analysis AnalysisResult `json:"analysis"`
}

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

type SentimentDistribution struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

type InfluenceScore struct {
	UserId         string `json:"user_id"`
	InfluenceScore int    `json:"influence_score"`
}

type Flags struct {
	MbrasEmployee      bool `json:"mbras_employee"`
	SpecialPattern     bool `json:"special_pattern"`
	CandidateAwareness bool `json:"candidate_awareness"`
}
