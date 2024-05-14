package first_optimized

var b = [20]byte{}

func squashRE(p string) []byte {
	pp := 0
	bp := 0
	for pp < len(p) {
		if pp+1 < len(p) && p[pp] == '.' && p[pp+1] == '*' {
			if bp > 1 && b[bp-1] == '*' {
				bp -= 2
			}
			b[bp] = p[pp]
			b[bp+1] = p[pp+1]
			bp += 2
			pp += 2
		} else if pp+1 < len(p) && p[pp] != '.' && p[pp+1] == '*' {
			if !(bp > 1 && (b[bp-2] == '.' || b[bp-2] == p[pp]) && b[bp-1] == '*') {
				b[bp] = p[pp]
				b[bp+1] = p[pp+1]
				bp += 2
			}
			pp += 2
		} else {
			b[bp] = p[pp]
			pp++
			bp++
		}
	}
	return b[:bp]
}

type rePart struct {
	t       int
	v       byte
	greed   int
	matched int
	pos     int
}

type re struct {
	parts   []rePart
	current int
}

func (r *re) parse(p []byte) {
	r.parts = r.parts[:0]
	r.current = 0
	pp := 0
	for pp < len(p) {
		if pp+1 < len(p) && p[pp+1] == '*' {
			r.parts = append(r.parts, rePart{t: 0, v: p[pp], greed: -1})
			pp += 2
		} else {
			r.parts = append(r.parts, rePart{t: 1, v: p[pp], greed: -1})
			pp++
		}
	}
}

func (r *re) match(s string) bool {
	sc := 0
	for i := range r.parts {
		part := &r.parts[i]
		part.matched = 0
		part.greed = len(s)
		part.pos = -1
	}
MainLoop:
	for r.current < len(r.parts) {
		part := &r.parts[r.current]
		if part.pos == -1 {
			part.pos = sc
		}
		if sc < len(s) && part.t == 0 && (s[sc] == part.v || part.v == '.') && (part.matched < part.greed) {
			sc++
			part.matched += 1
			continue
		} else if part.t == 1 {
			if sc < len(s) && (s[sc] == part.v || part.v == '.') {
				part.pos = sc
				sc++
				part.matched = 1
			} else {
				for i := r.current - 1; i >= 0; i-- {
					part := &r.parts[i]
					if part.t == 0 && part.greed > 0 && part.matched > 0 {
						r.current = i
						part.matched = 0
						part.greed--
						sc = r.parts[i].pos
						for j := r.current + 1; j < len(r.parts); j++ {
							part := &r.parts[j]
							part.pos = -1
							part.greed = len(s)
							part.matched = 0
						}
						continue MainLoop
					}
				}
				break
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
	}
	if len(s) == matchedSum && allMatched {
		return true
	}

	return false
}

func newRE() *re {
	r := &re{}
	r.parts = make([]rePart, 0, 20)
	return r
}

var r = newRE()

func isMatch(s string, p string) bool {
	pb := squashRE(p)
	r.parse(pb)
	return r.match(s)
}
