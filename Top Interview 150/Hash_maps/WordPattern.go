package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	str := strings.Split(s, ` `)
	if len(str) != len(pattern) {
		return false
	} 
	hashMapS := make(map[string]string)
	hashMapT := make(map[string]string)

	for i:=0 ; i<len(pattern) ; i++ {
		val, existS := hashMapS[string(str[i])]
		if existS && val != string(pattern[i]) {
			fmt.Println("this", i, val, string(str[i]), string(pattern[i]))
			return false
		}
		val, existT := hashMapT[string(pattern[i])]
		if existT && val != string(str[i]) {
			fmt.Println("that", i, string(val), string(str[i]), string(pattern[i]))
			return false
		}
		hashMapS[string(str[i])] = string(pattern[i])
		hashMapT[string(pattern[i])] = string(str[i])
	}
	return  true
}

func main() {
	fmt.Println(wordPattern("aaa", "aa aa aa aa"))
}