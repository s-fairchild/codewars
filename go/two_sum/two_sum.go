// Kata Two Sum
// See https://www.codewars.com/kata/52c31f8e6605bcc646000082/go for more info
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(TwoSum(arr, 4))
}

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}
