package textAnalysis

import(
	"strings"
)

// Determines if the string is just gibberish or not
func Gibberish(text string) bool {
	keyboardPatterns := []string {
        "qwerty", "asdf", "zxcv", "hjkl",
        "123456", "abcdef", "uvwxyz",
    }

	lowerText := strings.ToLower(text)
	for _, pattern := range keyboardPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}

	return false
}