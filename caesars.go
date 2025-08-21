package main

import (
	"fmt"
)

var ALPHABET string = "abcdefghijklmnopqrstuvwxyz"

func CaesarEncrypt(text string, shift int) string {
	var returnVal string = ""
	shift = Abs(shift) % 26
	for _, ch := range text {
		tempChar := ch + rune(shift)
		returnVal += string(tempChar)
	}
	return returnVal
}

func CaesarDecrypt(text string) string {
	fmt.Println("goodbye world!")
	return ""
}

func Abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}

func main() {
	fmt.Println(CaesarDecrypt("Banana"))
	fmt.Println(CaesarEncrypt("fruit", 1))
}
