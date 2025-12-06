package main

import (
	"slices"
)

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


func removeDuplicates(nums []int) int {    
	slices.Sort(nums)

	if len(nums) == 0 {
		return 0
	}
	i := 0

	for j := 0; j<len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}