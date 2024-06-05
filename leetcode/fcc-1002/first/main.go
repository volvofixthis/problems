package first

func commonChars(words []string) []string {
	gm := map[byte]int{}
	for _, v := range words[0] {
		gm[byte(v)] += 1
	}

	for i := 1; i < len(words); i++ {
		w := words[i]
		gl := map[byte]int{}
		for _, v := range w {
			gl[byte(v)] += 1
		}
		for k, v := range gm {
			gm[k] = min(v, gl[k])
		}
	}

	result := []string{}
	for k, v := range gm {
		if v > 0 {
			for i := 0; i < v; i++ {
				result = append(result, string(k))
			}
		}
	}
	return result
}
