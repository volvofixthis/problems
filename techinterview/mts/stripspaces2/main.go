package main

func stripSpaces(a []rune) []rune {
	if len(a) == 0 {
		return []rune{}
	}
	if len(a) == 1 {
		return a
	}
	p := 0
	sr := rune(' ')
	for _, v := range a[1:] {
		if a[p] == sr && v == sr {
			continue
		}
		p++
		a[p] = v
	}
	return a[:p+1]
}
