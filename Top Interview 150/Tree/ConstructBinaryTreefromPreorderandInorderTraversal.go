package main

// type TreeNode struct {
// 	Val int
//    	Left *TreeNode
//     Right *TreeNode
// }

func cons(root *TreeNode) *TreeNode {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	root := preorder[0]
	for i := 0 ; i<len(inorder) ; i++ {
		if inorder[i] == root {
			
		}
	}
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	buildTree(preorder, inorder)
}