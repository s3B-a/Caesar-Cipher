package textAnalysis

import(
	"regexp"
	"strings"
	"unicode"
)

// HasReasonableWords determines if the words within the script are reasonable and make sense
// and returns a boolean if the wordCount exceeds a threshold of 70%
func HasReasonableWords(text string) bool {
	words := strings.Fields(text)
	reasonableWordCount := 0

	for _, word := range words {
		cleanWord := regexp.MustCompile(`[^\p{L}]`).ReplaceAllString(word, "")
		if len(cleanWord) == 0 {
			continue
		}

		if looksReasonable(cleanWord) {
			reasonableWordCount++
		}
	}

	threshold := float64(len(words)) * 0.7
	return float64(reasonableWordCount) >= threshold
}

// Helper function to determine the characters in a  string, these characters
// (vowels and consonants) are used in tandem to determine the words length
// as well as determine if the word makes sense in the context of a given string
func looksReasonable(text string) bool {
	if len(text) == 0 {
		return false
	}

	if len(text) >= 2 {
		return true
	}

	vowels := "aeiouAEIOUyY"
	vowelCount := 0
	consonantCount := 0

	for _, ch := range text {
		if unicode.IsLetter(ch) {
			if strings.ContainsRune(vowels, ch) {
				vowelCount++
			} else {
				consonantCount++
			}
		}
	}

	totalLetters := vowelCount + consonantCount
	if totalLetters == 0 {
		return false
	}

	if vowelCount == 0 && len(text) > 3 {
		return false
	}

	if consonantCount == 0 && len(text) > 2 {
		return false
	}

	return true
}