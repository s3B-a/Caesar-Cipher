package textAnalysis

import(
	"strings"
)

// HumanReadable determines if a piece of text can be read by a human or not
func HumanReadable(text string) bool {
	if len(strings.TrimSpace(text)) == 0 {
		return false
	}

	if !HasBasicSentenceStructure(text) {
		return false
	}

	if !HasReasonableWords(text) {
		return false
	}

	if Gibberish(text) {
		return false
	}

	return true
}