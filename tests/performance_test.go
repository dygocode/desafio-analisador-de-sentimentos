package perf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

// Message representa uma única mensagem no feed
type Message struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	Timestamp string   `json:"timestamp"`
	UserID    string   `json:"user_id"`
	Hashtags  []string `json:"hashtags,omitempty"`
	Reactions int      `json:"reactions,omitempty"`
	Shares    int      `json:"shares,omitempty"`
	Views     int      `json:"views,omitempty"`
}

// MessageRequest representa a requisição para análise de mensagens
type MessageRequest struct {
	Messages          []Message `json:"messages"`
	TimeWindowMinutes int       `json:"time_window_minutes"`
}

// apiBaseURL retorna a URL base da API
func apiBaseURL() string {
	if v := os.Getenv("API_BASE_URL"); v != "" {
		return v
	}
	return "http://localhost:8000"
}

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

// postAnalyze envia uma requisição POST para o endpoint /analyze-feed
func postAnalyze(payload MessageRequest) (*http.Response, error) {
	url := apiBaseURL() + "/analyze-feed"

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return httpClient.Do(req)
}

// genDataset gera um conjunto de dados de teste com n mensagens
func genDataset(n int) MessageRequest {
	now := time.Date(2025, 9, 10, 11, 0, 0, 0, time.UTC)
	msgs := make([]Message, 0, n)

	for i := 0; i < n; i++ {
		ts := now.
			Add(-time.Duration(i%30) * time.Minute).
			Add(-time.Duration(i%5) * time.Second).
			Format(time.RFC3339)

		content := "Adorei o novo produto!"
		if i%4 == 0 {
			content = "ruim"
		}

		msgs = append(msgs, Message{
			ID:        formatID(i),
			Content:   content,
			Timestamp: ts,
			UserID:    formatUser(i),
			Hashtags:  hashtags(i),
			Reactions: (i % 7) + 1,
			Shares:    i % 3,
			Views:     ((i % 25) + 1) * 10,
		})
	}

	return MessageRequest{
		Messages:          msgs,
		TimeWindowMinutes: 30,
	}
}

// formatID auxilia a formatação de IDs
func formatID(i int) string {
	return fmt.Sprintf("perf_%04d", i)
}

// formatUser auxilia a formatação de UserIDs
func formatUser(i int) string {
	return fmt.Sprintf("user_%03d", i%200)
}

// hashtags auxilia na geração de hashtags para mensagens
func hashtags(i int) []string {
	if i%10 == 0 {
		return []string{"#produto", "#teste"}
	}
	return []string{"#produto"}
}

// TestPerformanceUnder200ms verifica se a análise de 1000 mensagens é concluída em menos de 200ms
func TestPerformanceUnder200ms(t *testing.T) {
	if os.Getenv("RUN_PERF") != "1" {
		t.Skip("Set RUN_PERF=1 to enable performance test")
	}

	payload := genDataset(1000)

	start := time.Now()
	resp, err := postAnalyze(payload)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}

	if elapsed > 200*time.Millisecond {
		t.Fatalf("Took %s", elapsed)
	}
}
