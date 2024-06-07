package first

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	dm := map[string]struct{}{}
	for _, prefix := range dictionary {
		dm[prefix] = struct{}{}
	}
	a := strings.Split(sentence, " ")
	for i, word := range a {
		for j := 1; j < len(word); j++ {
			k := word[:j]
			if _, ok := dm[k]; ok {
				a[i] = k
				break
			}
		}
	}
	return strings.Join(a, " ")
}
