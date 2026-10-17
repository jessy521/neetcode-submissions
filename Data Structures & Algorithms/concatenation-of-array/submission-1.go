func getConcatenation(nums []int) []int {
    res := make([]int,len(nums)*2)

	for i:=0;i<len(nums);i++{
		res[i],res[len(nums)+i] = nums[i],nums[i]
	}

	return res
}
