package main

func containsNearbyDuplicate(nums []int, k int) bool {
	numToIndexMap := make(map[int]int)

	for i, num := range nums {
		if idx, found := numToIndexMap[num]; found {
			if i-idx <= k {
				return true
			}
		}
		numToIndexMap[num] = i
	}

	return false
}
