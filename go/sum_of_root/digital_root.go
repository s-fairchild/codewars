// Kata Sum of Digits / Digital Root
// https://www.codewars.com/kata/541c8630095125aba6000c00
package main

import (
	"fmt"
	"strconv"
)

// 16  -->  1 + 6 = 7
// 942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
// 132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
// 493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(942))
	fmt.Println(DigitalRoot(132189))
	fmt.Println(DigitalRoot(493193))
}

func DigitalRoot(n int) int {
	var (
		narr []int
		sum int
	)
	narr = NumToArr(n)
	sum = FindSum(narr)

	// If sum is greater than 10, repeat and find the root
	for sum > 9 {
		narr = NumToArr(sum)
		sum = FindSum(narr)
	}

	return sum
}

// NumToArr accepts a number and returns an array with each digit
// being an array element
func NumToArr(n int) []int {
	// Convert int to string
	nstr := fmt.Sprint(n)
	var narr []int
	var err error
	var ntmp int
	for _, i := range nstr {
		ntmp, err = strconv.Atoi(string(i))
		narr = append(narr, ntmp)
		if err != nil {
			fmt.Println(err)
		}
	}
	return narr
}

// FindSum accepts an array of integers
// and returns the sum of them
func FindSum(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}