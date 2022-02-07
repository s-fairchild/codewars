// Kata Find Unique Numbers
// See https://www.codewars.com/kata/585d7d5adb20cf33cb000235/go for more information
package main

import "fmt"

func main() {
	array := []float32{ 1.0, 1.0, 1.0, 2.0, 1.0, 1.0 }
	fmt.Println(FindUniq(array))
}

func FindUniq(arr []float32) float32 {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[1] != arr[i] {
			if i != len(arr) - 1 {
				if arr[i] == arr[i+1] {
					return arr[1]
				} else {
					return arr[i]
				}
			} else {
				if arr[i] == arr[i-1] {
					return arr[1]
				} else {
					return arr[i]
				}
			}
		}
	}
	return 0
} 