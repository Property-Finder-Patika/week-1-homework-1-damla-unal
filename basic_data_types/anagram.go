package main

import "fmt"

func main() {
	fmt.Println(isAnagram("act", "tac")) //true
	fmt.Println(isAnagram("acd", "tac")) //false
}

// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letter in different order.
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		// these strings are not anagram
		return false
	}

	if s1 == s2 {
		return false
	}

	// Create a map of rune and int
	countMapByRune := make(map[rune]int)
	for _, c := range s1 {
		countMapByRune[c]++
	}

	for _, c := range s2 {
		countMapByRune[c]++
	}

	// when the all number of characters in map are even, these strings are anagram
	for _, count := range countMapByRune {
		if count%2 != 0 {
			return false
		}

	}
	return true

}
