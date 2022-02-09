package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "(){}[]"
	fmt.Println(ValidBraces(s)) // true

	s = "[(])"
	fmt.Println(ValidBraces(s)) // false
	s = "([{)]}"
	fmt.Println(ValidBraces(s)) // false
	s = "([}{])"
	fmt.Println(ValidBraces(s)) // false
}

func ValidBraces(str string) bool {
	var b int
	for _, r := range str {
		if r == '(' || r == '{' || r == '[' {
			b++
		} else {
			b--
		}
		if b < 0 {
			return false
		}
	}
	openPar :=  strings.Count(str, "(")
	openCur := strings.Count(str, "{")
	openSq := strings.Count(str, "[")
	closePar :=  strings.Count(str, ")")
	closeCur := strings.Count(str, "}")
	closeSq := strings.Count(str, "]")
	//fmt.Println("openPar:", openPar, "openCur:", openCur, "openSq:", openSq)
	if openPar != closePar || openCur != closeCur || closeSq != openSq {
		return false
	}
	return b == 0
}
