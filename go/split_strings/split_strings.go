// Kata Split Strings
// See https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/go for more information
package main

import "fmt"

func main() {
	s := "awsaws"
	fmt.Println(Solution(s))
	s = "abc"
	fmt.Println(Solution(s))
	s = "dawsd"
	fmt.Println(Solution(s))
}

func Solution(str string) []string {
	// If the string length is not even, add an underscore
	// This is for the kata requirements, and so the for loop doesn't go out of the array index range
	if (len(str)) % 2 != 0 {
		str += "_"
	}
	sSlice := make([]string, 0)
	n := 2
	// Split the string into two rune strings
	for i := 0; i < len(str); i+=2 {
		if n <= 2 {
			tmpstr := string(str[i])
			tmpstr += string(str[i+1])
			sSlice = append(sSlice, tmpstr)
			n = 0
		}
		n++
	}
	return sSlice
}