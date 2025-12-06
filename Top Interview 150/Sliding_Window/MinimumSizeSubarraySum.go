package main

func minSubArrayLen(target int, nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	var minLength int
	hasMinLength := 0
	sum := 0
	left := 0

	for right := 0 ; right < len(nums) ; right++ {
		sum += nums[right]

		for sum >= target {
			lenSum := right - left + 1
			if lenSum < minLength || minLength == 0 {
				minLength = lenSum
			}
			hasMinLength = 1
			sum -= nums[left]
			left++
		}
	}
	if hasMinLength == 1{
		return minLength
	} else {
		return 0
	}
}

// func main() {
// 	nums := []int{1,1,1,1,1,1,1,1}
// 	fmt.Println(minSubArrayLen(11, nums))
// }