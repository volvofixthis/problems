package first

import (
	"strings"
)

var results = []string{}
var wordSet = map[string]struct{}{}

func cleanWordSet() {
	for k := range wordSet {
		delete(wordSet, k)
	}
}

func wordBreak(s string, wordDict []string) []string {
	cleanWordSet()
	for _, v := range wordDict {
		wordSet[v] = struct{}{}
	}
	results = results[:0]
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
