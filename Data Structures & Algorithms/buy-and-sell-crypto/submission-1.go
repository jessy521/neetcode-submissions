func maxProfit(prices []int) int {
	if len(prices) == 0{return 0}
	
	profit := 0
	minPrice := prices[0]

	for _,cp := range prices{
		if cp < minPrice{
			minPrice = cp
		}else if  cp - minPrice > profit{
			profit = cp - minPrice
		}
	}

	return profit
}
