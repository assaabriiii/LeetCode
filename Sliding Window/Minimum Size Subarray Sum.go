package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1

	for idx, _ := range nums {
		sum += nums[idx]
		for sum >= target {
			if minLen > idx-left+1 {
				minLen = idx - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
