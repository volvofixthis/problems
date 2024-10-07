package first

import (
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{"/"}
	for _, v := range strings.Split(path, "/") {
		switch v {
		case ".", "":
			continue
		case "..":
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)
		}

	}
	return filepath.Join(stack...)
}
