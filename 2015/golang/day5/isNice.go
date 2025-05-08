package main

import (
	"fmt"
)

var Count_nice = 0

func IsNice(line string) {

	fmt.Printf("La linea es: %v \n", line)
	// palabras prohibidas
	forbidenCheck := ForbidenMatch(line)
	// vocales presentes
	vowelCheck := VowelsPresent(line)
	// letras repetidas
	repetitionCheck := RepeatedLetters(line)
	// contando si es nice word
	if vowelCheck == true && forbidenCheck == true && repetitionCheck == true {
		Count_nice += 1
	}

	fmt.Printf("Nice words: %v\n", Count_nice)
}
