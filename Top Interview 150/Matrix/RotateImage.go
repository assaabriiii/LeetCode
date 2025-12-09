package main

import "fmt"

func reverseInts(input []int) []int {
    if len(input) == 0 {
        return input
    }
    return append(reverseInts(input[1:]), input[0]) 
}

func rotate(matrix [][]int)  {
	fmt.Println(matrix)

	for i:=0 ; i<len(matrix); i++ {
		for j:= i+1; j<len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		} 
	}
	fmt.Println(matrix)

	for i:=0 ; i<len(matrix) ; i++ {
		matrix[i] = reverseInts(matrix[i])
	}
	fmt.Println(matrix)
}

// func main() {
// 	matrix := [][]int{
// 		{1,2,3},
// 		{4,5,6},
// 		{7,8,9},
// 	}
// 	rotate(matrix)
// }