package main

import "fmt"

// func isIsomorphic(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	hashMapS := make(map[string]string)
// 	hashMapT := make(map[string]string)

// 	for i:=0 ; i<len(s) ; i++ {
// 		val, existS := hashMapS[string(s[i])]
// 		if existS && val != string(t[i]) {
// 			return false
// 		}
// 		val, existT := hashMapT[string(t[i])]
// 		if existT && val != string(s[i]) {
// 			return false
// 		}
// 		hashMapS[string(s[i])] = string(t[i])
// 		hashMapT[string(t[i])] = string(s[i])
// 	}
// 	return  true
// }


func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[string]string)

	for i:=0 ; i < len(s) ; i++ {
		if val, exist := hashMap[string(s[i])]; exist {
			if val != string(t[i]) {
				fmt.Println("that?", i)
				return false
			}
		}
		if val, exist := hashMap[string(t[i])]; exist {
			if val == string(t[i]) {
				fmt.Println("this?", i, val , string(t[i]))
				return false
			}
		}
		hashMap[string(s[i])] = string(t[i])
	}
	fmt.Println(hashMap)
	return true
}

// func main() {
// 	fmt.Println(isIsomorphic("badc", "baba")) // b-b, a-a, d-b, c-a "paper", "title" p-t, a-i, p-t, e-l, r-e
// 	fmt.Println(isIsomorphic("paper", "title"))  // b
// 	fmt.Println(isIsomorphic("egcd", "adfd")) // e-a, g-d, c-f, d-d
// }
