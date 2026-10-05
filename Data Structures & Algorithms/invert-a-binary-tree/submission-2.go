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

	qu := []*TreeNode{root}

	for len(qu)>0{
		curr := qu[0]
		qu = qu[1:]

		curr.Left, curr.Right = curr.Right, curr.Left

		if curr.Left != nil{
			qu = append(qu,curr.Left)
		}
		if curr.Right != nil{
			qu = append(qu,curr.Right)
		}
	}

	return root
}
