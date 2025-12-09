package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	dir := "right"
	ans := make([]int, m*n)

	i:=0
	j:=0
	k:=0

	for k < m*n {
		ans[k] = matrix[i][j]
		switch dir {
		case "right":
			if i+j == n-1 {
				dir = "down"
				i++
			} else {
				j++
			}
		case "down":
			if n-j == m-i {
				dir = "left"
				j--
			} else {
				i++
			}
		case "left":
			if i+j == m-1 {
				dir = "up"
				i--
			} else {
				j--
			}
		case "up":
			if i-j == 1 {
				dir = "right"
				j++
			} else {
				i--
			}
		}
		k++
	}
	return ans
}
// func main() {
// 	matrix := [][]int{
// 		{1,2,3,10},
// 		{4,5,6,20}, 
//  		{7,8,9,30}, 
// 		// [0][0] 1 -> [0][1] 2 -> [0][2] 3 -> [0][3] 10 --> i++ --> [1][3] end of j increase i
// 		// [1][3] 20 -> [2][3] 30 --> j-- --> [2][2] end of i and end of j decrease j
// 		// [2][2] 9 -> [2][1] 8 -> [2][0] 7 --> i-- --> [1][0] 
// 		// -> [1][0] 4 -> [1][1] 5 -> [1][2] 6
// 	}				

// 	fmt.Println(spiralOrder(matrix))
// }