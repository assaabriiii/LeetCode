package main

func isHappy(n int) bool {
	results := make(map[int]bool)

	for n != 1 && !results[n] {
		results[n] = true
		res := 0

		for n > 0 {
			digit := n % 10
			res += digit * digit
			n /= 10
		}
		n = res
	}
	return n == 1
}
