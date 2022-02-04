// Kata Create Phone Number
// https://www.codewars.com/kata/525f50e3b73515a6db000b83/go
package main

import (
	"fmt"
)

func main() {
	nums := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	phoneNum := CreatePhoneNumber(nums) // returns "(123) 456-7890"
	fmt.Println(phoneNum)
}

// CreatePhoneNumber accepts an array of uints of size 10
// and returns those numbers as a string formatted as a phone number
func CreatePhoneNumber(numbers [10]uint) string {
	phoneNumber := "("
	for i, n := range numbers {
		phoneNumber += fmt.Sprint(n)
		if i == 2 {
			phoneNumber += ") "
		} else if i == 5 {
			phoneNumber += "-"
		}
	}
	return phoneNumber
}
