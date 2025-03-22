package main

import "strings"

func canConstruct(ransomNote string, magazine string) bool {
	for _, letter := range ransomNote {
		if strings.Count(ransomNote, string(letter)) > strings.Count(magazine, string(letter)) {
			return false
		}
	}
	return true
}

// understand this one

// func canConstruct(ransomNote string, magazine string) bool {

// 	freq := make([]int, 26)

// 	for _, r := range magazine {
// 		freq[r-'a']++
// 	}

// 	for _, r := range ransomNote {
// 		freq[r-'a']--
//         if freq[r-'a'] < 0 {
//             return false
//         }
// 	}

// 	return true

// }
