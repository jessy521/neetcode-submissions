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
    res := []int{}

	inorder(root,&res)

	// sort.Slice(res, func(j, k int) bool {
	// 	return res[j] < res[k]
	// })

	// for _,v :=range res{fmt.Println(v)}

	return res[k-1]
}
