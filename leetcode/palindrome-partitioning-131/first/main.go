package first

var result = [][]string{}

func isPalindrome(s string) bool {
	l := len(s)
	m := l / 2
	if m == 0 {
		return true
	}
	for i := 0; i <= m; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func backtrack(s string, path []string, result *[][]string) {
	if len(s) == 0 {
		*result = append(*result, append([]string(nil), path...))
		return
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			backtrack(s[i:], append(path, s[:i]), result)
		}
	}
}

func partition(s string) [][]string {
	result = result[:0]
	backtrack(s, []string{}, &result)
	return result
}
