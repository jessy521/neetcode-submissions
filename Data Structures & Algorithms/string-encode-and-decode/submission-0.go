type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0{
		return ""
	}

	var sizes []string
	for _, str := range strs{
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return strings.Join(sizes,",")+ "#" + strings.Join(strs,"")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == ""{
		return []string{}
	}
	parts := strings.SplitN(encoded,"#",2)
	sizes := strings.Split(parts[0],",")

	var res []string
	count := 0
	for i:=0;i<len(sizes);i++{
		size := sizes[i]
		if size == ""{
			continue
		}
		realSize,_ := strconv.Atoi(size)
		len := count+realSize
		res = append(res,parts[1][count:len])
		count = len
	}
	return res
}
