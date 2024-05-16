package first

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return false
	}
	m := len(s) / 2
	for i := 0; i <= m; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	return ""
}
