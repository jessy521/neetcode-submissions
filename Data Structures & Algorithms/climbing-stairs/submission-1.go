func memo(n int,res []int)int{
	if n<=2{
		return n
	}

	if res[n] != 0{
		return res[n]
	}
	
	res[n] = memo(n-1,res) + memo(n-2,res)

	return res[n]
}

func climbStairs(n int) int {
    res := make([]int,n+1)
	return memo(n, res)
}
