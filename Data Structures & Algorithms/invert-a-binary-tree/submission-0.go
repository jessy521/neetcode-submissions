/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}

    temp := root

	l := invertTree(temp.Left)
	r := invertTree(temp.Right)

	root.Right = l
	root.Left = r

	return root
}
