package first

import (
	"strings"
)

func removeChar(s string, i int) string {
	return s[:i] + s[i+1:]
}

func clearStars(s string) string {
	sma := []int{}
	p := 0
	sm := byte('z') + 1
	for p < len(s) {
		if s[p] == '*' {
			sl := removeChar(s[:p+1], sma[0])
			slc := 0
			for i := 1; i < len(sma); i++ {
				sc := removeChar(s[:p+1], sma[i])
				if strings.Compare(sc, sl) == -1 {
					sl = sc
					slc = i
				}
			}
			s = sl[:p-1] + s[p+1:]
			p--
			sma = append(sma[:slc], sma[slc+1:]...)
			for i := slc; i < len(sma); i++ {
				sma[i] -= 1
			}
			if len(sma) == 0 {
				sm = 'z' + 1
				for i := 0; i < p; i++ {
					if sm > s[i] {
						sm = s[i]
						sma = []int{i}
					} else if sm == s[i] {
						sma = append(sma, i)
					}
				}
			}
		} else {
			if sm > s[p] {
				sm = s[p]
				sma = []int{p}
			} else if sm == s[p] {
				sma = append(sma, p)
			}
			p++
		}
	}
	return s
}
