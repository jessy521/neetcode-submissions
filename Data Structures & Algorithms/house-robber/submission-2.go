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
	n := len(nums)

    // return dfs(0,nums,res)

	if n==0{return 0}
	if n == 1{return nums[0]}

	res := make([]int,n)
	res[0] = nums[0]
	res[1] = max(nums[0],nums[1])

	for i:=2;i<n;i++{
		res[i] = max(res[i-1] , nums[i]+res[i-2])
	}

	return res[n-1]
}
