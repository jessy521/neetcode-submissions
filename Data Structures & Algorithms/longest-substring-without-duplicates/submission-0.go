func lengthOfLongestSubstring(s string) int {
	charset := make(map[byte]bool)
	l,maxLen := 0,0

	for r:=0 ; r<len(s); r++{
		for charset[s[r]]{
			delete(charset,s[l])
			l++
		}
		charset[s[r]] = true

		if r-l+1 > maxLen {
			maxLen = r-l+1
		}
	}

	return maxLen
}
