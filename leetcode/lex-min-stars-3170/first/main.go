package first

import (
	"strings"
)

func removeChar(s string, i int) string {
	return s[:i] + s[i+1:]
}

func clearStars(s string) string {
	p := 1
	for p < len(s) {
		if s[p] == '*' {
			sm := s[0]
			smp := 0
			for i := 0; i < p; i++ {
				if sm > s[i] {
					sm = s[i]
					smp = i
				}
			}
			sl := removeChar(s, smp)
			for i := smp; i < p; i++ {
				if s[i] != sm {
					continue
				}
				sc := removeChar(s, i)
				if strings.Compare(sc, sl) == -1 {
					sl = sc
				}
			}
			s = sl
			p--
			s = removeChar(s, p)
		} else {
			p++
		}
	}
	return s
}
