package main

import "fmt"


func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	max := 0

	j := len(height) - 1 

	for i:=0 ; i<len(height);{
		fmt.Println(i, j)
		if i == j {
			break
		}
		if height[i] > height[j] {
			area := height[j] * (j-i)
			if area > max {
				max = area
			}
			j--
		} else {
			area := height[i] * (j-i)
			if area > max {
				max = area
			}
			i++
		}
	}
	return max
}

func main() {
	// 2
	// 3
	// 
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}