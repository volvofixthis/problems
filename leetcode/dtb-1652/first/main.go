package main

func sumNext(code []int, k int, i int) int {
	r := 0
	for k > 0 {
		k--
		i++
		if i > len(code)-1 {
			i = 0
		}
		r += code[i]
	}
	return r
}

func sumPrev(code []int, k int, i int) int {
	r := 0
	for k > 0 {
		k--
		i--
		if i < 0 {
			i = len(code) - 1
		}
		r += code[i]
	}
	return r
}

func decrypt(code []int, k int) []int {
	r := make([]int, len(code))
	if k > 0 {
		for i := 0; i < len(code); i++ {
			r[i] = sumNext(code, k, i)
		}
	} else if k < 0 {
		k = k * -1
		for i := 0; i < len(code); i++ {
			r[i] = sumPrev(code, k, i)
		}
	}
	return r
}
