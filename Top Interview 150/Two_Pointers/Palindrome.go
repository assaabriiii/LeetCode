package main

import (
	"regexp"
	"strings"
)


func removeRegex(in string) string {
	in = strings.ToLower(in)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]*")
	return reg.ReplaceAllString(in, "")
}

// al ec : alecs 

func isPalindrome(s string) bool {
	s = removeRegex(s) // delete

	if len(s) == 1 {
		return true
	}

	j := len(s) - 1
	for i:=0 ; i<len(s)/2 ; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
