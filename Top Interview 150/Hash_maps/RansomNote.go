package main

import "fmt"


func canConstruct(ransomNote string, magazine string) bool {
	hash := make(map[string]int)

	for i:=0 ; i<len(magazine) ; i++ {
		hash[string(magazine[i])] += 1
	}

	for i:=0 ; i<len(ransomNote) ; i++ {
		if _, exists := hash[string(ransomNote[i])]; !exists {
			return false
		} else {
			hash[string(ransomNote[i])] -=1
			if hash[string(ransomNote[i])] < 0 {
				return false
			}
		}
	}
	return  true 
}

func main() {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
}