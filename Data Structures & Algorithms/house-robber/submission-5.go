
func rob(nums []int) int {
	n := len(nums)

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
