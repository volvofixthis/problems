package first

func appendCharacters(s string, t string) int {
	sp := 0
	tp := 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[tp] {
			tp += 1
		}
		sp += 1
	}
	return len(t) - tp
}
