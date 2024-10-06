package first

func checkIfPangram(sentence string) bool {
	a := map[rune]struct{}{}
	for _, l := range sentence {
		a[l] = struct{}{}
	}
	return len(a) == 26
}
