package main

func stripSpaces(a []rune) []rune {
	if len(a) == 0 {
		return []rune{}
	}
	if len(a) == 1 {
		return a
	}
	r := []rune{a[0]}
	sr := rune(' ')
	for _, v := range a[1:] {
		if r[len(r)-1] == sr && v == sr {
			continue
		}
		r = append(r, v)
	}
	return r
}
