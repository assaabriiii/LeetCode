package main

import "fmt"


func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return nil
	}

	for j:= 0 ; j<len(numbers)-1 ; j++ {
		for i:=j+1 ; i<len(numbers) ; i++ {
			if numbers[i] + numbers[j] == target {
				return []int{j+1, i+1}
			}
		}
	}
	return nil
}

func main(){
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums, 9))
}