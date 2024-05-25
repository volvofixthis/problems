package first

var max_score = 0
var freq = map[int]int{}

func isValidWord(subset_letters map[int]int) bool {
	for i := 0; i < 26; i++ {
		if freq[i] < subset_letters[i] {
			return false
		}
	}
	return true
}

func check(w int, words []string, score []int, subset_letters map[int]int, total_score int) {
	if w == -1 {
		max_score = max(max_score, total_score)
		return
	}
	check(w-1, words, score, subset_letters, total_score)
	L := len(words[w])
	for i := 0; i < L; i++ {
		c := int(words[w][i] - 97)
		subset_letters[c] += 1
		total_score += score[c]
	}

	if isValidWord(subset_letters) {
		check(w-1, words, score, subset_letters, total_score)
	}

	for i := 0; i < L; i++ {
		c := int(words[w][i] - 97)
		subset_letters[c] -= 1
		total_score -= score[c]
	}
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	W := len(words)
	max_score = 0
	freq = map[int]int{}
	subset_letters := map[int]int{}
	for _, c := range letters {
		freq[int(c-97)] += 1
	}

	check(W-1, words, score, subset_letters, 0)
	return max_score
}
