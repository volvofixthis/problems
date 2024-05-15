package first

type ce map[byte]int

var c = make(ce, 50)

func resetC() {
	for k := range c {
		delete(c, k)
	}
}

func lengthOfLongestSubstring(s string) int {
	ml := 0
	p := 0
	resetC()
	for i := 0; i < len(s); i++ {
		v := s[i]
		e, ok := c[v]
		if ok && p <= e {
			p = e + 1
		}
		c[v] = i
		cl := i - p + 1
		if cl > ml {
			ml = cl
		}
	}
	return ml
}
