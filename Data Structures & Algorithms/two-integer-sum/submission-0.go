func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)
	for i, n := range nums {
        indices[n] = i
    }

	for i,j := range nums{
		diff := target - j
		if x,found := indices[diff]; found && i!=x {
			return []int{i, x}
		}
	}
	return []int{}

}
