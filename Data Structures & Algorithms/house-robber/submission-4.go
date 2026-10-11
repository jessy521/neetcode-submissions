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
	// res:=make([]int,n)

    // return dfs(n-1,nums,res)

	prev := nums[0]
	prev2 := 0

	for i:= 1;i<n;i++ {
		take := nums[i]
		if i>1 {take += prev2}
		noTake := prev
		 
		cur := max(take,noTake)
		prev2 = prev
		prev = cur
	}

	return prev
}
