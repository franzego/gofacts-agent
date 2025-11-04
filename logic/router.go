package logic

import (
	"strings"
)

func contains(msg string, keywords []string) bool {
	for _, k := range keywords {
		if strings.Contains(msg, k) {
			return true
		}
	}
	return false
}

func DetectCategory(logic string) string {

	msg := strings.ToLower(strings.TrimSpace(logic))
	switch {
	case contains(msg, []string{"tip", "advice", "best practice"}):
		return "💡 " + TipsRandom()
	case contains(msg, []string{"fact", "fun", "go fun"}):
		return "🤓 " + FactRandom()
	case contains(msg, []string{"code", "example", "syntax", "how to"}):
		return "💻 " + SyntaxRandom()
	case contains(msg, []string{"history", "origin", "gopher"}):
		return "📜 " + HistoryRandom()
	default:
		return "Hey! I can share Go tips, facts, syntax, or fun trivia. Try asking for one! 🐹"
	}
}
