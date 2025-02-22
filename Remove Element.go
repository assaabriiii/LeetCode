package main

func removeElement(nums []int, val int) int {
	result := []int{}

	for _, value := range nums {
		if value != val {
			result = append(result, value)
		}
	}
	return len(result)
}
