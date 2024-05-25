package first

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := map[string]struct{}{}
	for _, v := range wordDict {
		wordSet[v] = struct{}{}
	}
	results := []string{}
	backtrack(s, wordSet, []string{}, &results, 0)
	return results
}

func backtrack(s string,
	wordSet map[string]struct{},
	currentSentence []string,
	results *[]string,
	startIndex int,
) {
	if startIndex == len(s) {
		*results = append(*results, strings.Join(currentSentence, " "))
		return
	}

	for i := startIndex; i < len(s)+1; i++ {
		word := s[startIndex:i]
		if _, ok := wordSet[word]; ok {
			backtrack(s, wordSet, append(currentSentence, word), results, i)
		}
	}
}
