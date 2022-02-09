// Kata WeIrD StRiNg CaSe
// See https://www.codewars.com/kata/52b757663a95b11b3d00062d/go for more info
package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "abc def"
	fmt.Println(toWeirdCase(s)) // Answer "AbC DeF"
	s = "ABC"
	fmt.Println(toWeirdCase(s)) // Answer "AbC"
	s = "This is a test Looks like you passed"
	fmt.Println(toWeirdCase(s)) // Answer "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"
}

func toWeirdCase(str string) string {
	strArr := strings.Split(str, " ")
	var (
		strResult string
		n         int
	)
	for _, word := range strArr {
		for i := 0; i < utf8.RuneCountInString(word); i++ {
			if i%2 == 0 {
				strResult += string(unicode.ToUpper(rune(word[i])))
			} else {
				strResult += string(unicode.ToLower(rune(word[i])))
			}
		}
		if n != len(strArr)-1 {
			strResult += " "
		}
	}
	return strings.Trim(strResult, " ")
}
