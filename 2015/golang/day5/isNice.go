package main

import (
	"fmt"
)

var Count_nice = 0
var Count_nice_second = 0

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
		//fmt.Printf("%v\n", line)
	}

	// 426 muy alto
	// 174 es muy bajo
	// 255 right answer
	// fmt.Printf("Nice words: %v\n", Count_nice)

	// Second star
	// ver que tengan letras repetidas con un letra en medio
	betweenLettersCheck := BetweenLetters(line)

	// verificar que se repita un patron de dos letras
	patternCheck := Pattern(line)

	if betweenLettersCheck && patternCheck {
		Count_nice_second += 1
		fmt.Printf("LINE: %v\n", line)
		fmt.Printf("Count: %v\n", Count_nice_second)
	}

	// 381
	// 7 -> Not the right answer
	// 39 -> not the right answer
	// 53 -> not the right answer
	// 55 -> the right answer
}
