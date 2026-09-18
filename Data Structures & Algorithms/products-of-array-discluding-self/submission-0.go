func productExceptSelf(nums []int) []int {
	allProduct := 1
	zeroCount :=0

	for _,v := range nums{
		if v!=0{
			allProduct *= v
		}else{
			zeroCount++
		}
	}

	res := make([]int,len(nums))
	if zeroCount > 1{
		return res
	}
	for i,v := range nums{
		if zeroCount > 0{
			if v == 0{
				res[i] = allProduct
			}else{
				res[i] = 0
			}
		}else{
			res[i] = allProduct/v 
		}
	}

	return res
}
