/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil{
		return 0
	}

	l1 := maxDepth(root.Left)
	l2 := maxDepth(root.Right)

	return max(l1,l2)+1
}
