package main

import (
	"fmt"
)

var Count_nice = 0

func IsNice(line string) {

	// palabras prohibidas
	forbidenCheck := ForbidenMatch(line)
	// vocales presentes
	vowelCheck := VowelsPresent(line)
	// letras repetidas
	repetitionCheck := RepeatedLetters(line)
	// contando si es nice word
	if vowelCheck == true && forbidenCheck == true && repetitionCheck == true {
		Count_nice += 1
		fmt.Printf("%v\n", line)
	}

	// 426 muy alto
	// 174 es muy bajo
	// 255 right answer
	fmt.Printf("Nice words: %v\n", Count_nice)
}
