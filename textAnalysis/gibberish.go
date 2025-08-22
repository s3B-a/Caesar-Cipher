package textAnalysis

import (
	"strings"
)

// Determines if the string is just gibberish or not
func Gibberish(text string) bool {
	keyboardPatterns := []string{
		"qwerty", "asdf", "zxcv", "hjkl",
		"123456", "abcdef", "uvwxyz",
	}

	// Checks for common keyboard patterns that might be gibberish
	lowerText := strings.ToLower(text)
	for _, pattern := range keyboardPatterns {
		if strings.Contains(lowerText, pattern) {
			return true
		}
	}

	// Detect long runs of non-words
	if len(lowerText) > 20 && !strings.Contains(lowerText, " ") {
		return true
	}

	// Detect repeated characters (e.g., "aaaaaa", "zzzzzz")
	if len(lowerText) > 5 {
		repeated := true
		for i := 1; i < len(lowerText); i++ {
			if lowerText[i] != lowerText[0] {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}

	vowels := "aeiou"
	vowelCount := 0
	for _, ch := range lowerText {
		if strings.ContainsRune(vowels, ch) {
			vowelCount++
		}
	}
	if len(lowerText) > 10 && vowelCount < 2 {
		return true
	}

	if len(lowerText) > 6 {
		isAlt := true
		for i := 1; i < len(lowerText); i++ {
			if (isLetter(lowerText[i-1]) && isLetter(lowerText[i])) || (isDigit(lowerText[i-1]) && isDigit(lowerText[i])) {
				isAlt = false
				break
			}
		}
		if isAlt {
			return true
		}
	}

	return false
}

// Helper to check if a byte is a letter by taking in a byte and outputting
// a boolean value: true if the byte corresponds to an ASCII letter [a-z][A-Z], false otherwise
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// Helper to check if a byte is a digit by taking in a byte and outputting
// a boolean value: true if the byte corresponds to an ASCII digit [0-9], false otherwise
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
