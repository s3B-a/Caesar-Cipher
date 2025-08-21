package textAnalysis

import(
	"strings"
	"unicode"
)

// Checks if a scentence is written like a traditional scentence
func HasBasicSentenceStructure(text string) bool {
	hasLetter := false
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return false
	}

	words := strings.Fields(text)
	if len(words) < 1 {
		return false
	}

	validSingleChars := map[string]bool{"I": true, "a": true, "A": true}
	if len(words) == 1 && len(words[0]) == 1 {
		return validSingleChars[words[0]]
	}

	return true
}