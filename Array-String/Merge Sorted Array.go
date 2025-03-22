package main


import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	merged := append(nums1[:m], nums2[:n]...)
	sort.Ints(merged)
	return merged
}