// Kata Longest vowel chain
// See https://www.codewars.com/kata/59c5f4e9d751df43cf000035/go for more info
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "codewarriors"
	fmt.Println(Solve(s))

	s = "suoidea"
	fmt.Println(Solve(s))
}

// Solve accepts a string and returns the count of the longest vowel substring
func Solve(s string) int {
	regex := regexp.MustCompile(`[aeiou]+`) // Regex search pattern for vowels, + means prefer higher rune matches
	matches := regex.FindAllString(s, -1) // Save matches
	if len(matches) > 0 { // Find the longest if matches were found
		var longestMatchIndex int
		for i := range matches {
			if len(matches[i]) >= len(matches[longestMatchIndex]) {
				longestMatchIndex = i
			}
		}
		fmt.Println("Index:", longestMatchIndex)
		return len(matches[longestMatchIndex])
	} else {
		fmt.Println("Error no matches found, matches has length of:", matches)
	}
	return 0
}
