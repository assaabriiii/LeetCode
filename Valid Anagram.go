package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqs := make(map[rune]int)

	for _, r := range s {
		freqs[r]++
	}

	letter := len(freqs)
	for _, r := range t {
		_, exists := freqs[r]
		if !exists {
			return false
		}
		freqs[r]--
		if freqs[r] < 0 {
			return false
		}
		if freqs[r] == 0 {
			letter--
		}
	}
	return letter == 0
}
