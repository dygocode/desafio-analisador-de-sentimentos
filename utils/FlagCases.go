package utils

import (
	"regexp"
	"time"
)

var (
	mbrasUserPattern = regexp.MustCompile(`mbras`)
	candidatePattern = regexp.MustCompile(`teste técnico mbras`)
	specialPattern   = regexp.MustCompile(`mbras`)
)

func IsMBRASUser(userId string) bool {
	normalizedUserId := NormalizeForMatching(userId)
	return mbrasUserPattern.MatchString(normalizedUserId)
}

func IsCandidateAwareness(content string) bool {
	return candidatePattern.MatchString(content)
}

func IsSpecialPattern(content string) bool {
	if len(content) < 42 {
		return false
	}
	return specialPattern.MatchString(content)
}

func IsTimeWindow(
	timestamp string,
	timeWindowMinutes int,
) bool {

	now := time.Now().UTC()
	layout := "2006-01-02T15:04:05Z"
	// Parse do timestamp da mensagem
	minutesMessage, err := time.Parse(layout, timestamp)
	if err != nil {
		return false
	}

	// 2️⃣ Calcula o limite inferior
	lowerBound := now.Add(-time.Duration(timeWindowMinutes))

	// 3️⃣ Verifica se o timestamp está dentro da janela
	return minutesMessage.Equal(lowerBound) || minutesMessage.After(lowerBound)
}
