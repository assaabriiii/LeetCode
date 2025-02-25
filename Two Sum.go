package main


func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if idx, found := numToIndex[diff]; found {
			return []int{i, idx}
		}
		numToIndex[num] = i
	}
	return nil
}