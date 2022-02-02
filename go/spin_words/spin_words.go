package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	result := SpinWords("Hey fellow warriors")
	fmt.Println(result)
}

func SpinWords(str string) string {
	rs := strings.Split(str, " ")
	spunstr := make([]string, 0)
	for i := 0; i < len(rs); i++ {
		if utf8.RuneCountInString(rs[i]) >= 5 {
			tmpstr := reverse(rs[i])
			spunstr = append(spunstr, tmpstr)
		} else {
			spunstr = append(spunstr, rs[i])
		}
	}
	var finalstr string
	for i, s := range spunstr {
		finalstr += s
		if i < len(spunstr) - 1 {
			finalstr += " "
		}
	}
	return finalstr
}

// Reverse accepts a string and returns that string
// reveresed
func reverse(str string) string {
	rs := []rune(str)
	var nw string
	for j := len(rs) - 1; j >= 0; j-- {
		nw += string(rs[j])
	}
	return nw
}
