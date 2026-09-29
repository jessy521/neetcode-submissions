func characterReplacement(s string, k int) int {
count := make(map[byte]int)
res,l,maxF := 0,0,0

for r:=0;r<len(s);r++{
	count[s[r]]++
	if count[s[r]] > maxF{
		maxF = count[s[r]]
	}

	if r-l+1 - maxF > k{
		count[s[l]]--
		l++
	}

	if r-l+1 > res{
		res = r-l+1
	}
}

return res
}
