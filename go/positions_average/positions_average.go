// Kata Positions Average
// See https://www.codewars.com/kata/59f4a0acbee84576800000af/go
package main

import (
	"fmt"
	"strings"
	//"math"
)

func main() {
	s := "466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096" // correct 26.6666666667
	fmt.Println(PosAverage(s))

	s = "444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490" // correct 29.2592592593
	fmt.Println(PosAverage(s))

	s = "6900690040, 4690606946, 9990494604"
	fmt.Println(PosAverage(s))
}

func PosAverage(s string) float64 {
	sArr := strings.Split(s, ", ")
	var (
		n float64
		comparisonCount float64
	)
	for i := 0; i < len(sArr); i++ {
		for j := i+1; j < len(sArr); j++ {
			for k := 0; k < len(sArr[i]); k++ {
				comparisonCount++
				if sArr[i][k] == sArr[j][k] {
					n++
				}
			}
		}
	}
	return 100.0 * (n / comparisonCount)
}