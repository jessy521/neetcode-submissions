func lengthOfLongestSubstring(s string) int {
	charset := make(map[byte]int)
	l,maxLen := 0,0

	for r:=0 ; r<len(s); r++{
		if idx, found := charset[s[r]]; found{
			l = max(idx+1,l)
		}
		charset[s[r]] = r

		if r-l+1 > maxLen {
			maxLen = r-l+1
		}
	}

	return maxLen
}
