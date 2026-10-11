func dfs(i int,nums,res []int)int{
	// if i>=len(nums){
	// 	return 0
	// }

	if i == 0{
		return nums[0]
	}
	if i < 0{
		return 0
	}

	if res[i]>0{
		return res[i]
	}

	pick := nums[i] + dfs(i-2,nums,res)
	notPick := dfs(i-1,nums,res)
	res[i] = max(pick,notPick )

	return res[i]
}

func rob(nums []int) int {
	n := len(nums)
	res:=make([]int,n)

    return dfs(n-1,nums,res)

	// if n==0{return 0}
	// if n == 1{return nums[0]}

	// res := make([]int,n)
	// res[0] = nums[0]
	// res[1] = max(nums[0],nums[1])

	// for i:=2;i<n;i++{
	// 	res[i] = max(res[i-1] , nums[i]+res[i-2])
	// }

	// return res[n-1]
}
