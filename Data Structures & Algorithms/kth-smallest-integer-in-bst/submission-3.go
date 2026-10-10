/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(root *TreeNode, res *[]int){
	if root == nil{
		return
	}

	inorder(root.Left,res)
	*res = append(*res, root.Val)
	inorder(root.Right,res)
}

func kthSmallest(root *TreeNode, k int) int {
	curr := root

	for { 
		if curr.Left == nil{
			k--
			if k == 0 {
				return curr.Val
			}
			curr = curr.Right
		}else{
			pred := curr.Left
			for pred.Right != nil && pred.Right != curr{
				pred = pred.Right
			}
			if pred.Right == nil{
				pred.Right = curr
				curr = curr.Left
			}else{
				pred.Right = nil
				k--
				if k == 0{
					return curr.Val
				}
				curr = curr.Right
			}
		}
	}
}
