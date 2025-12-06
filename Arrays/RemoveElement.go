package main

import "fmt"

// Source - https://stackoverflow.com/a/37335777
// Posted by T. Claverie, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-02, License - CC BY-SA 4.0

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


func removeElement(nums []int, val int) (int, []int) {
	if len(nums) == 0 {
		return 0, nums
	}
	
	for i := 0; i<len(nums); i++ {
		if nums[i] == val {
			nums = remove(nums, i)
			i--
		}
	}
	return len(nums), nums
}

// func main() {
// 	nums := []int{0,1,2,2,3,0,4,2}
// 	val := 2
// 	fmt.Println(removeElement(nums, val))
// }