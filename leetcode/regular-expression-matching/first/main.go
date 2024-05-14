package first

func squashRE(p string) string {
	result := ""
	pp := 0
	for pp < len(p) {
		if pp+1 < len(p) && p[pp] == '.' && p[pp+1] == '*' {
			if len(result) >= 2 && result[len(result)-1] == '*' {
				result = result[:len(result)-2]
			}
			result = result + string(p[pp]) + string(p[pp+1])
			pp += 2
		} else if pp+1 < len(p) && p[pp] != '.' && p[pp+1] == '*' {
			if !(len(result) >= 2 && (result[len(result)-2] == '.' || result[len(result)-2] == p[pp]) && result[len(result)-1] == '*') {
				result = result + string(p[pp]) + string(p[pp+1])
			}
			pp += 2
		} else {
			result = result + string(p[pp])
			pp++
		}
	}
	return result
}

type rePart struct {
	t       int
	v       byte
	greed   int
	matched int
	s       string
	pos     int
}

type re struct {
	parts   []*rePart
	current int
}

func (r *re) match(s string) bool {
	sc := 0
	for _, part := range r.parts {
		part.matched = 0
		part.greed = len(s)
		part.pos = -1
	}
	for r.current < len(r.parts) {
		part := r.parts[r.current]
		if part.pos == -1 {
			part.pos = sc
		}
		if sc < len(s) && part.t == 0 && (s[sc] == part.v || part.v == '.') && (part.matched < part.greed) {
			part.s = part.s + string(s[sc])
			sc++
			part.matched += 1
			continue
		} else if part.t == 1 {
			if sc < len(s) && (s[sc] == part.v || part.v == '.') {
				part.pos = sc
				part.s = string(s[sc])
				sc++
				part.matched = 1
			} else {
				// fmt.Println("trying to rewind from: ", r.current)
				oldCurrent := r.current
				for i := r.current - 1; i >= 0; i-- {
					part := r.parts[i]
					if part.t == 0 && part.greed > 0 && part.matched > 0 {
						r.current = i
						// fmt.Println("rewinded to: ", r.current)
						part.s = ""
						part.matched = 0
						part.greed--
						sc = part.pos
						for j := r.current + 1; j < len(r.parts); j++ {
							part := r.parts[j]
							part.pos = -1
							part.s = ""
							part.greed = len(s)
							part.matched = 0
						}
						break
					}
				}
				if oldCurrent == r.current {
					break
				}
				continue
			}
		}
		r.current++
	}
	allMatched := true
	matchedSum := 0
	for _, part := range r.parts {
		matchedSum += part.matched
		if part.matched == 0 && part.t == 1 {
			allMatched = false
		}
		// if part.t == 0 {
		// 	fmt.Println("wildcard part: .", string(part.v), ", matched: ", part.s, ", pos: ", part.pos)
		// } else {
		// 	fmt.Println("symbol part: ", string(part.v), ", matched: ", part.s, ", pos: ", part.pos)
		// }
	}
	if len(s) == matchedSum && allMatched {
		return true
	}

	return false
}

func newRE(p string) *re {
	parts := []*rePart{}
	pp := 0
	for pp < len(p) {
		if pp+1 < len(p) && p[pp+1] == '*' {
			parts = append(parts, &rePart{t: 0, v: p[pp], greed: -1})
			pp += 2
		} else {
			parts = append(parts, &rePart{t: 1, v: p[pp], greed: -1})
			pp++
		}
	}
	return &re{parts: parts}
}

func isMatch(s string, p string) bool {
	p = squashRE(p)
	r := newRE(p)
	return r.match(s)
}
