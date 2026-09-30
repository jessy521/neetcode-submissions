func minWindow(s string, t string) string {
    m,n := len(t),len(s)
	seen := [256]int{}
	l,count := 0,0
	maxi, si := math.MaxInt64,-1

	for _,i := range t{
		seen[i]++
	}

	for r:=0;r<n;r++{
		if seen[s[r]]>0{
			count++
		}
		seen[s[r]]--

		for count == m{
			if r-l+1 <maxi{
				maxi = r-l+1
				si = l
			}
			seen[s[l]]++
			if seen[s[l]] > 0{
				count--
			}
			l++
		}

	}
	if si == -1{
		return ""
	}

	return s[si : si+maxi]
}
