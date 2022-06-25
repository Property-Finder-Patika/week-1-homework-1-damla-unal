package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findBaseName("a/b/c.go")) // "c"
	fmt.Println(findBaseName("c.d.go"))   // "c.d"
	fmt.Println(findBaseName("abc"))      // "abc"

	fmt.Println(findBaseNameWithStringsLastIndex("a/b/c.go")) // "c"
	fmt.Println(findBaseNameWithStringsLastIndex("c.d.go"))   // "c.d"
	fmt.Println(findBaseNameWithStringsLastIndex("abc"))      // "abc"

}

// findBaseName removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func findBaseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			i++
			path = path[i:]
			break
		}
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}
	return path

}

func findBaseNameWithStringsLastIndex(path string) string {
	// -1 if "/" not found
	lastSlash := strings.LastIndex(path, "/")
	path = path[lastSlash+1:]
	if lastDot := strings.LastIndex(path, "."); lastDot >= 0 {
		path = path[:lastDot]
	}
	return path

}
