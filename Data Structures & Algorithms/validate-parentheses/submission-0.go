type Stack struct{
	items []rune
}

func (s *Stack)Push(el rune){
	s.items = append(s.items, el)
}

func (s *Stack)IsEmpty()bool{
	if len(s.items) == 0{
		return true
	}
	return false
}

func (s *Stack)Pop(){
    if !s.IsEmpty() {
        s.items = s.items[:len(s.items)-1]
    }
}

func (s *Stack)Top()rune {
    if s.IsEmpty() {
        return 0
    }
    return s.items[len(s.items)-1]
}

func isValid(s string) bool {
	val := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	if len(s)%2 != 0{
		return false
	}
    st := Stack{}

	for _, ch := range s {
		if _, ok := val[ch]; ok {
			st.Push(ch)
		} else {
			if st.IsEmpty() || val[st.Top()] != ch {
				return false
			}
			st.Pop()
		}
	}
	return st.IsEmpty()
}
