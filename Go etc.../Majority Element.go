package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	// Iterate through the array
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
		fmt.Println(count)
	}

	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority element:", majorityElement(nums))
}
