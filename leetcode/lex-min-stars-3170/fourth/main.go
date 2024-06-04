package first

func less(s1 []byte, s1k int, s2 []byte, s2k int) bool {
	s1p := 0
	s2p := 0
	s1s := 0
	s2s := 0
	equal := true
	for s1p < len(s1) || s2p < len(s2) {
		// fmt.Println(s1p, s2p)
		if s1p >= len(s1) {
			s2p++
		} else if s2p >= len(s2) {
			s1p++
		} else if s1[s1p] == ' ' || s1k == s1p {
			// fmt.Println("skipping s1")
			s1s += 1
			s1p++
		} else if s2[s2p] == ' ' || s2k == s2p {
			// fmt.Println("skipping s2")
			s2s += 1
			s2p++
		} else if s1[s1p] > s2[s2p] {
			// fmt.Println("symbol from s1p is bigger")
			return false
		} else if s1[s1p] < s2[s2p] {
			equal = false
			break
		} else {
			s1p++
			s2p++
		}
	}
	if equal && len(s1)-s1s > len(s2)-s2s {
		return false
	}
	return true
}

func clearStars(s string) string {
	sa := []byte(s)
	cp := map[byte][]int{}
	p := 0
	sm := byte('z') + 1
	for p < len(sa) {
		if s[p] == '*' {
			sma := cp[sm]
			slc := 0
			sl := sa[sma[slc]:p]
			for i := 1; i < len(sma); i++ {
				if less(sl, sma[i]-sma[0], sl, sma[slc]-sma[0]) {
					slc = i
				}
			}
			sa[sma[slc]] = ' '
			sa[p] = ' '
			cp[sm] = cp[sm][:len(cp[sm])-1]
			copy(cp[sm][slc:], sma[slc+1:])
			if len(cp[sm]) == 0 {
				delete(cp, sm)
				sm = 'z' + 1
				for k := range cp {
					if sm > k {
						sm = k
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
	r := ""
	for _, v := range sa {
		if v != ' ' {
			r += string(v)
		}
	}
	return r
}
