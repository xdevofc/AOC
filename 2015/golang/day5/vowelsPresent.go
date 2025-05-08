package main

import (
	"strings"
)

var count_vowels = 0

func VowelsPresent(line string) bool {

	for _, vowel := range Vowels {
		if strings.Contains(line, vowel) {
			count_vowels += 1
		}
	}

	if count_vowels >= 3 {
		return true
	} else {
		count_vowels = 0
		return false
	}
}
