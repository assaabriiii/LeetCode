package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	
	for i:=0 ; i<len(nums) ; i++ {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]] += 1
		} else {
			m[nums[i]] = 1
		}
	}
	for key, val := range m {
		if val > len(nums)/2 {
			return key
		}
	}
	return 0
}

func main() {
	arr := []int{1,1,1,1,1,2,2,2,2,2,3}
	fmt.Println(majorityElement(arr))
}