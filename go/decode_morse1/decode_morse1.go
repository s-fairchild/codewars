// Kata Decode the Morse code part 1
// https://www.codewars.com/kata/54b724efac3d5402db00065e/go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. .")) // = HEY JUDE
	fmt.Println(DecodeMorse("   . .")) // =  E E
}

func DecodeMorse(morseCode string) string {
	var MORSE_CODE = map[string]string{
		".-": "A",
		"-...": "B",
		"-.-.": "C",
		"-..": "D",
		".": "E",
		"..-.": "F",
		"--.": "G",
		"....": "H",
		"..": "I",
		".---": "J",
		"-.-": "K",
		".-..": "L",
		"--": "M",
		"-.": "N",
		"---": "O",
		".--.": "P",
		"--.-": "Q",
		".-.": "R",
		"...": "S",
		"-": "T",
		"..-": "U",
		"...-": "V",
		".--": "W",
		"-..-": "X",
		"-.--": "Y",
		"--..": "Z",
		"...---...": "SOS",
		".-.-.-": ".",
		"--..--": ",",
		"..--..": "?",
		"-.-.--": "!",
	}

	// Split morse code into an array of words
	ws := strings.Split(morseCode, "   ")

	var words string
	for i := 0; i < len(ws); i++ {
		// Split each word into array index
		word := strings.Split(ws[i], " ")
		for _, w := range word {
			// Use MorseCode MAP to convert to English chars
			// Build string of chars
			words += MORSE_CODE[w]
		}
		words += " "
	}
	return strings.Trim(words, " ") // Return without leading or trailing whitespace
}
