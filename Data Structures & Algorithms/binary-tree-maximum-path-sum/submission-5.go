/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    // res := -1 << 31
    // dfs(root, &res)
    // return res

	res := []int{root.Val}

	var dfs func(node *TreeNode)int

	dfs = func(node *TreeNode)int{
		if node == nil{
			return 0
		}
		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		res[0] = max(res[0], node.Val + leftMax + rightMax)
		return node.Val + max(leftMax , rightMax)
	}

	dfs(root)

	return res[0]
}

func dfs(root *TreeNode, res *int) {
    if root == nil {
        return
    }
    left := getMax(root.Left)
    right := getMax(root.Right)
    *res = max(*res, root.Val + left + right)
    dfs(root.Left, res)
    dfs(root.Right, res)
}

func getMax(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := getMax(root.Left)
    right := getMax(root.Right)
    path := root.Val + max(left, right)
    return max(0, path)
}