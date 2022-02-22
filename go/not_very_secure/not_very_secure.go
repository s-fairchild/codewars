// Kata Not Very Secure
// See https://www.codewars.com/kata/526dbd6c8c0eb53254000110/go for more information
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "43534h56jmTHHF3k"
	fmt.Println(alphanumeric(s))
}

func alphanumeric(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z-0-9]+$`).MatchString(str)
}
