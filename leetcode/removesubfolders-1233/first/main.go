package main

import (
	"sort"
)

func IsSubfolder(sub, root string) bool {
	if len(sub) < len(root) {
		return false
	}
	end := len(root)
	for _, v := range sub[end:] {
		if v == '/' {
			break
		}
		end++
	}
	if sub[:end] == root {
		return true
	}
	return false
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	p := 0
	for i := 0; i < len(folder); i++ {
		if IsSubfolder(folder[i], folder[p]) {
			continue
		}
		p++
		folder[p] = folder[i]
	}
	return folder[:p+1]
}
