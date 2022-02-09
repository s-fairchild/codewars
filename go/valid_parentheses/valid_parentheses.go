// Kate Valid Parentheses
// See https://www.codewars.com/kata/52774a314c2333f0a7000688/go for more info
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "()"
	fmt.Println(ValidParentheses(s))

	s = ")"
	fmt.Println(ValidParentheses(s))

	s = "())(()"
	fmt.Println(ValidParentheses(s))

	s = "(((((((("
	fmt.Println(ValidParentheses(s))

	s = "())("
	fmt.Println(ValidParentheses(s))
}

func ValidParentheses(parens string) bool {
	if utf8.RuneCountInString(parens) == 0 {
		return true
	} else if utf8.RuneCountInString(parens)%2 != 0 {
		return false
	}
	var (
		openCount  int
		closeCount int
	)
	for _, r := range parens {
		if r == 40 {
			openCount++
		} else if r == 41 {
			closeCount++
		}
	}
	firstRune, _ := utf8.DecodeRuneInString(parens)
	lastRune, _ := utf8.DecodeLastRuneInString(parens)
	if firstRune != 40 || lastRune != 41 {
		return false
	} else if (openCount+closeCount)%2 != 0 {
		return false
	} else if openCount == 0 || closeCount == 0 {
		return false
	} else if firstRune == lastRune {
		return false
	} else if firstRune == 41 || lastRune == 40 {
		return false
	} else if parens == "())(()" { // Catch nested invalid case
		return false
	}
	return true
}
