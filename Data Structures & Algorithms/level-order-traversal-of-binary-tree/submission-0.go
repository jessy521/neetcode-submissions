/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	qu := []*TreeNode{root}
	var res [][]int
	if root == nil {return res}

	for len(qu) > 0{
		n := len(qu)
		var tmp []int
		for i:=0;i<n;i++{
			node := qu[0]
			tmp = append(tmp,node.Val)
			if node.Left != nil{qu = append(qu, node.Left)}
			if node.Right != nil{qu = append(qu, node.Right)}
			qu = qu[1:]
		}
		res = append(res,tmp)
	}
	return res
}
