package main

func OneEditApart(a, b string) bool {
	ap := 0
	bp := 0
	r := 0
	var inc func()
	if len(a) == len(b) {
		inc = func() {
			ap++
			bp++
		}
	} else if len(a) < len(b) {
		inc = func() {
			bp++
		}
	} else if len(a) > len(b) {
		inc = func() {
			ap++
		}
	}
	for ap < len(a) && bp < len(b) {
		if a[ap] != b[bp] {
			inc()
			r++
		} else {
			ap++
			bp++
		}
	}
	for bp < len(b) {
		bp++
		r++
	}
	for ap < len(a) {
		ap++
		r++
	}
	return r <= 1
}
