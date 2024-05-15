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
	cl := 0
	resetC()
	for i := 0; i < len(s); i++ {
		v := s[i]
		e, _ := c[v]
		if e != i {
			cl++
		} else {
			cl = i - (e + 1)
		}
		c[v] = i
		if cl > ml {
			ml = cl
		}
	}
	return ml
}
