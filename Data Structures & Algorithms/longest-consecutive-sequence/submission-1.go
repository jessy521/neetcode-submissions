func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})

	for _,num := range nums{
		numSet[num] = struct{}{}
	}

	longset := 0
	for num := range numSet{
		if _, found := numSet[num-1]; !found{
			length := 1
			for{
				if _,exist := numSet[num+length];exist{
					length++
				}else{
					break
				}
			}
			if length>longset{
				longset = length
			}
		}
	}

	return longset
}
