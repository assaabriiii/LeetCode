package main




func setZeroes(matrix [][]int)  {
	cols := len(matrix[0])
	rows := len(matrix)
	frz, fcz := false, false

	for j:=0 ; j<cols ; j++ {
		if matrix[0][j] == 0 {
			fcz = true
			break
		}
	}
	for i:=0 ; i<rows ; i++ {
		if matrix[i][0] == 0 {
			frz = true
			break
		}
	}

	for i:= 1 ; i<rows ; i++ {
		for j:=1 ; j<cols ; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	
}

// func main() {
// 	matrix := [][]int{
// 		{0,1,2,0},{3,4,5,2},{1,3,1,5},
// 	}
// 	setZeroes(matrix)
// }