// Steven Fairchild
// Kata Completed kata Find The Parity Outlier
// https://www.codewars.com/kata/5526fc09a1bbd946250002dc

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 6, 8, -10, 3}
	outlier := FindOutlier(arr)
	fmt.Println(outlier)

	arr = []int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}
	outlier = FindOutlier(arr)
	fmt.Println(outlier)

	arr = []int{17, 6, 8, 10, 6, 12, 24, 36}
	outlier = FindOutlier(arr)
	fmt.Println(outlier)
}

func FindOutlier(integers []int) int {
	var (
		odd_count  int
		even_count int
	)
	// Find out if there's more even's or odd's
	for _, i := range integers {
		if i%2 == 0 {
			even_count++
		} else {
			odd_count++
		}
		if even_count > 1 || odd_count > 1 {
			break
		}
	}
	fmt.Println("odd count: ", odd_count, "even count: ", even_count)
	// If there are more even numbers in the array
	// search for the one odd number
	if even_count > 1 {
		for _, i := range integers {
			if i%2 != 0 || i == 1 {
				return i
			}
		}
		// If there's more odd numbers
		// then search for the one even number
	} else {
		for _, i := range integers {
			if i%2 == 0 {
				return i
			}
		}
	}
	return 0
}
