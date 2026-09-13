func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _,s := range strs{
		sorted := sortString(s)
		res[sorted] = append(res[sorted],s)
	}

	var result [][]string

	for _, grp := range res{
		result = append(result, grp)
	}

	return result
}

func sortString(s string)string{
	charecters := []rune(s)

	sort.Slice(charecters, func(i,j int)bool{
		return charecters[i]>charecters[j]
	})

	return string(charecters)
}
