package main


func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []int{}

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            currentLevel = append(currentLevel, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, currentLevel)
    }

    return result
}


func averageOfLevels(root *TreeNode) []float64 {
	res := levelOrder(root)
	result := []float64{}

	for i := 0 ; i<len(res) ; i++ {
		r := 0
		for j := 0 ; j<len(res[i]) ; j++ {
			r += res[i][j]
		}
		result = append(result, r)
	}
	return result
}