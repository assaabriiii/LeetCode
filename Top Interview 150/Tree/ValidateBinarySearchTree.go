package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val > root.Right.Val {
		return false
	}
	if root.Val < root.Left.Val {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}