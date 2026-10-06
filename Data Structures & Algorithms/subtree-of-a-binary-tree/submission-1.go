/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isSame(root *TreeNode,subRoot *TreeNode)bool{
	if root == nil && subRoot == nil{return true}

	if root != nil && subRoot != nil && root.Val == subRoot.Val{
		return isSame(root.Left,subRoot.Left) && isSame(root.Right,subRoot.Right)
	}

	return false
 }

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil{return false}
	if subRoot == nil{return true}

	if isSame(root,subRoot){return true}

	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}
