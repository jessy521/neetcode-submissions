func memo(nums []int)int{
	if len(nums) == 0{return 0}

	if len(nums) == 1{return nums[0]}

	// dp := make([]int,len(nums))
	prev := nums[0]
	prev2 := 0

	for i:=1;i<len(nums);i++{
		pick := nums[i]
		if i>1{pick += prev2}
		notPick := prev

		curr := max(pick,notPick)
		prev2 = prev
		prev = curr
	}
	return prev
}

func rob(nums []int) int {
	if len(nums) == 1{return nums[0]}

	return max( memo(nums[1:]), memo(nums[:len(nums)-1]))
}
