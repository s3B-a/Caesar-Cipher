# 🛡️ Caesar Cipher ASCII Encryption Tool

This project provides a flexible Caesar Cipher implementation written in Go that can encrypt and decrypt a wide range of ASCII characters, including:

- Lowercase letters (a–z)  
- Uppercase letters (A–Z)  
- Digits (0–9)  
- Most other characters are passed through unchanged

What sets this implementation apart is its intelligent decryption system that attempts to recover the original text without a known shift value, using a built-in text analysis engine to determine if the output is human-readable.

---

## 🔧 Features

### Caesar Encryption

- Encrypts letters and digits using the Caesar Cipher method  
- Accepts any positive or negative integer shift value  
- Normalizes shifts to a standard range for consistent behavior
**Example:**
```
  CaesarEncrypt("FRUIT16324a", 898751)
  // Output: "KYZNZ61879f"
```
---

### Caesar Decryption

- Attempts all 26 possible Caesar shifts  
- Uses text analysis to detect the most likely correct decryption  
- Automatically selects the most human-readable result
**Example:**
```
CaesarDecrypt("67g8161zt42fsty32f2hu359")
// Output: "FRUIT16324a"
```
---

## 🧠 Human-Readable Text Detection

The text analysis engine evaluates whether a string is likely to be human-readable based on multiple heuristics:

- Presence of sentence-like structure  
- Recognition of reasonable word patterns  
- Avoidance of gibberish (keyboard mashing, repeated characters, no vowels, etc.)

This logic is critical for decrypting text when the original Caesar shift is unknown.

---

## 🗂️ Project Structure

caesar-cipher/  
├── main.go - Core Caesar encryption and decryption logic  
└── textAnalysis/ - Text evaluation utilities  
  ├── human.go - Determines if a string is human-readable  
  ├── sentence.go - Checks for basic sentence structure  
  ├── words.go - Evaluates the reasonableness of words  
  └── gibberish.go - Detects gibberish or nonsensical strings

---

## 📦 Usage

To run the program, use the Go CLI:

go run main.go

This will output the result of an encrypted string and its automatically detected decryption.
