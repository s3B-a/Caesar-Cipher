package main

import (
	"caesar-cipher/textAnalysis"
	"fmt"
)

// Caesar Cipher Encryption that allows for the majority of ascii characters
// to be encrypted in a Caesar Cipher algorithm
// [a-z][A-Z][0-9] and more
func CaesarEncrypt(text string, shift int) string {
	var returnVal string = ""
	shift = Abs(shift) % 95
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			pos := int(ch - 'a')
			newPos := (pos + shift) % 26
			returnVal += string('a' + newPos)
		} else if ch >= 'A' && ch <= 'Z' {
			pos := int(ch - 'A')
			newPos := (pos + shift) % 26
			returnVal += string('A' + newPos)
		} else if ch >= '0' && ch <= '9' {
			pos := int(ch - '0')
			newPos := (pos + shift) % 10
			returnVal += string('0' + newPos)
		} else {
			returnVal += string(ch)
		}
	}
	return returnVal
}

// Caesar Cipher Deryption that allows for the majority of ascii characters
// to be decrypted from a given Caesar Cipher
// [a-z][A-Z][0-9] and more
func CaesarDecrypt(text string) string {
	for shift := 0; shift < 26; shift++ {
		var returnVal string = ""

		for _, ch := range text {
			if ch >= 'a' && ch <= 'z' {
				pos := int(ch - 'a')
				newPos := (pos - shift + 26) % 26
				returnVal += string('a' + newPos)
			} else if ch >= 'A' && ch <= 'Z' {
				pos := int(ch - 'A')
				newPos := (pos - shift + 26) % 26
				returnVal += string('A' + newPos)
			} else if ch >= '0' && ch <= '9' {
				pos := int(ch - '0')
				newPos := (pos - shift + 10) % 10
				returnVal += string('0' + newPos)
			} else {
				returnVal += string(ch)
			}
		}

		if textAnalysis.HumanReadable(returnVal) {
			return returnVal
		}
	}

	return "No readable decryption found"
}

// Abs (Absolute Value) returns the absolute value of an integer
func Abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}

func main() {
	fmt.Println(CaesarDecrypt("67g8161zt42fsty32f2hu359"))
	fmt.Println(CaesarEncrypt("FRUIT16324a", 898751))
}
