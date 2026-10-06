/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	type Pair struct{
		first,second *TreeNode
	}

	qu := []Pair{{p,q}}

	for len(qu) > 0{
		node1,node2 := qu[0].first, qu[0].second
		qu = qu[1:]

		if node1 == nil && node2 == nil{continue}

		if node1 == nil || node2 == nil || node1.Val != node2.Val{
			return false
		}

		qu = append(qu,Pair{node1.Left,node2.Left})
		qu = append(qu,Pair{node1.Right,node2.Right})
	}
	return true
}