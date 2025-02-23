package main

func removeDuplicates(nums []int) int {
	seen := make(map[int]bool)
	result := []int{}

	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return len(result)
}
