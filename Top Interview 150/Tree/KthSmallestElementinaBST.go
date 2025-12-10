package main

func kthSmallest(root *TreeNode, k int) int {
    var result int
    count := 0

    var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }

        inorder(node.Left)

        count++
        if count == k {
            result = node.Val
            return  // early exit (optional, but faster)
        }

        inorder(node.Right)
    }

    inorder(root)
    return result
}