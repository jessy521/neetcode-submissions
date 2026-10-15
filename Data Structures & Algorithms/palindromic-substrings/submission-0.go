func countSubstrings(s string) int {
	oddCount, evenCount := 0,0
	for i:=0;i<len(s);i++{
		l,r := i,i
		for l>=0 && r<len(s) && s[l] == s[r]{
			oddCount++
			l--
			r++
		}

		l,r = i,i+1
		for l>=0 && r<len(s) && s[l] == s[r]{
			evenCount++
			l--
			r++
		}
	}
	return oddCount + evenCount
}
