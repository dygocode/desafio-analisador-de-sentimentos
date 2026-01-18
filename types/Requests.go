package types

// Estruturas do Request
type MessageRequest struct {
	Messages          []Message `json:"messages" validate:"required,min=1,dive"`
	TimeWindowMinutes int       `json:"time_window_minutes" validate:"required,min=1"`
}

type Message struct {
	Id        string `json:"id" validate:"required"`
	Content   string `json:"content" validate:"required,max=280"`
	Timestamp string `json:"timestamp" validate:"required"`
	UserId    string `json:"user_id" validate:"required"`

	Hashtags []string `json:"hashtags,omitempty"`

	Reactions int `json:"reactions,omitempty" validate:"gte=0"`
	Shares    int `json:"shares,omitempty" validate:"gte=0"`
	Views     int `json:"views,omitempty" validate:"gte=0"`
}
