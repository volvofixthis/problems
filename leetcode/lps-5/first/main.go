package first

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
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
	max := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i <= j && s[j] == s[i] {
				sub := s[i : j+1]
				if isPalindrome(sub) {
					if len(sub) > len(max) {
						max = sub
					}
				}
			}

		}
	}
	return max
}
