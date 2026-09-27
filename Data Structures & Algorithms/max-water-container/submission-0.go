func maxArea(heights []int) int {
	res := 0
	j,k := 0 , len(heights)-1


	for j<k {
		vol := (k-j) * min(heights[k],heights[j])
		if vol > res{
			res = vol
		}
		if heights[j]<heights[k]{
			j++
		}else{
			k--
		}
	}

	return res
}
