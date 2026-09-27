func maxProfit(prices []int) int {
	profit := 0;

	for i:=0;i<len(prices)-1;i++{
		buy := prices[i]
		for j := i+1;j <len(prices);j++{
			sell := prices[j]
			if sell<buy{
				continue
			}
			temp := sell - buy
			if temp>profit{
				profit = temp
			}
		}
	}
	return profit
}
