package main

import (
	"strings"
	"unicode"
)

func reverseString(s string) string {

	if len(s) <= 1 {
		return s
	}

	return reverseString(s[1:]) + string(s[0])
}

func removeNonAlphanumeric(s string) string {
	result := []rune{}

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result = append(result, char)
		}
	}

	return string(result)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)
	reversed := reverseString(s)
	if reversed == s {
		return true
	}
	return false
}
