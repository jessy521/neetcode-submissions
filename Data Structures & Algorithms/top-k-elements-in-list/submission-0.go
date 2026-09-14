func topKFrequent(nums []int, k int) []int {
	res := make(map[int]int)

	for _,v := range nums{
		res[v]++
	}

	keys := make([]int,0,len(res))
	for k := range res{
		keys = append(keys,k)
	}

	sort.Slice(keys,func(i,j int)bool{
		return res[keys[i]] > res[keys[j]]
	})

	return keys[:k]
}
