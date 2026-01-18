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

// Função para verificar se o userId pertence a um usuário MBRAS
func IsMBRASUser(userId string) bool {
	normalizedUserId := NormalizeForMatching(userId)
	return mbrasUserPattern.MatchString(normalizedUserId)
}

// Função para verificar se o conteúdo da mensagem indica candidato ciente do teste
func IsCandidateAwareness(content string) bool {
	return candidatePattern.MatchString(content)
}

// Função para verificar se o conteúdo da mensagem segue um padrão especial
func IsSpecialPattern(content string) bool {
	if len(content) < 42 {
		return false
	}
	return specialPattern.MatchString(content)
}

// Função para verificar se o timestamp está dentro da janela temporal especificada
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

	lowerBound := now.Add(-time.Duration(timeWindowMinutes))

	return minutesMessage.Equal(lowerBound) || minutesMessage.After(lowerBound)
}
