package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(nums []int) int {    
	slices.Sort(nums)

	if len(nums) == 0 {
		return 0
	}

	i := 1
	c := 1

	for j := 1; j<len(nums); j++ {
		if nums[j] == nums[j-1] {
			c += 1
		} else {
			c = 1
		}
		if c <= 2 {
			nums[i] = nums[j]
			i++
		}
		fmt.Println(c)
	}
	return i
}

func main() {
	nums := []int{1,1,1,1,1,2,2,2,2,2,3,4,5,6}
	fmt.Println(removeDuplicates(nums))
}