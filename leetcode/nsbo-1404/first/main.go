package first

func divideByTwo(sa []byte) []byte {
	return sa[:len(sa)-1]
}

func addOne(sa []byte) []byte {
	i := len(sa) - 1
	for i >= 0 && sa[i] != '0' {
		sa[i] = '0'
		i -= 1
	}

	if i < 0 {
		sa = append([]byte{'1'}, sa...)
	} else {
		sa[i] = '1'
	}
	return sa
}

func numSteps(s string) int {
	sa := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		sa[i] = s[i]
	}

	operations := 0
	for len(sa) > 1 {
		if sa[len(sa)-1] == '0' {
			sa = divideByTwo(sa)
		} else {
			sa = addOne(sa)
		}
		operations += 1
	}
	return operations
}
