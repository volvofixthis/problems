package first

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Println(s, p)
	pp := 0
	sc := 0
	greed := -1
	arg := ""
	matchPp := 0
	matchSc := 0
	matchLen := 0
	for {
		fmt.Println("sc pp", sc, pp)
		if pp+1 < len(p) && p[pp+1] == '*' {
			fmt.Println("Greed", greed, string(p[pp]))
			if (sc < len(s) && (p[pp] == s[sc] || p[pp] == '.')) && (greed == -1 || len(arg) < greed) {
				if arg == "" {
					matchSc = sc
					matchPp = pp
				}
				arg = arg + string(s[sc])
				matchLen = len(arg)
				fmt.Println(arg)
				sc++
			} else {
				pp += 2
			}
		} else if sc < len(s) && (p[pp] == '.' || s[sc] == p[pp]) {
			sc++
			pp++
		} else if matchLen > 0 {
			fmt.Println("decreasing greed")
			greed = matchLen - 1
			matchLen = 0
			sc = matchSc
			arg = ""
			pp = matchPp
		} else {
			return false
		}
		if pp >= len(p) {
			break
		}
	}
	return sc >= len(s)
}
