package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLength := 0
	left := 0
	hashSet := make(map[rune]bool)

	for right := 0; right < len(s); right++ {
		for hashSet[rune(s[right])] {
			delete(hashSet, rune(s[left]))
			left++
		}
		hashSet[rune(s[right])] = true
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

// func main() {
// 	a := "pwwkew"
// 	fmt.Println(lengthOfLongestSubstring(a))
// }