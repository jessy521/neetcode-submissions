func dfs(i int,nums,res []int)int{
	if i>=len(nums){
		return 0
	}

	if res[i]>0{
		return res[i]
	}

	res[i] = max(dfs(i+1,nums,res), nums[i] + dfs(i+2,nums,res))
	return res[i]
}
func rob(nums []int) int {
	res := make([]int,len(nums)+1)
    return dfs(0,nums,res)
}
