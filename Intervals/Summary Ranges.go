// not finished
package main

import "strconv"

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}

	for idx := range len(nums) - 1 {
		if nums[idx]-nums[idx+1] > 1 {
			result = append(result, strconv.Itoa(nums[idx]))
		}
	}
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
}

// Input: nums = [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
// Explanation: The ranges are:
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"
