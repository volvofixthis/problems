package first

import (
	"strings"
	"unsafe"
)

var sa = []byte{}

func fillSA(s string) {
	l := len(s)
	d := unsafe.StringData(s)
	b := unsafe.Slice(d, l)
	if cap(sa) < l {
		sa = make([]byte, l)
	}
	sa = sa[:l]
	copy(sa, b)
}

var cp = map[byte][]int{}

func clearCP() {
	for k := range cp {
		cp[k] = cp[k][:0]
	}
}

func clearStars(s string) string {
	fillSA(s)
	clearCP()
	p := 0
	sm := byte('z')
	for p < len(sa) {
		if s[p] == '*' {
			sma := cp[sm]
			slc := len(sma) - 1
			sa[sma[slc]] = ' '
			sa[p] = ' '
			cp[sm] = cp[sm][:len(cp[sm])-1]
			copy(cp[sm][slc:], sma[slc+1:])
			if len(cp[sm]) == 0 {
				fsm := sm
				sm = 'z'
				for i := fsm + 1; i <= 'z'; i++ {
					if len(cp[i]) > 0 {
						sm = i
						break
					}
				}
			}
		} else if s[p] != ' ' {
			if sm >= s[p] {
				sm = s[p]
			}
			cp[s[p]] = append(cp[s[p]], p)
		}
		p++
	}
	var b strings.Builder
	for _, v := range sa {
		if v != ' ' {
			b.WriteByte(v)
		}
	}
	return b.String()
}
