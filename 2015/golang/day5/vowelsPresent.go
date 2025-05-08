package main

import (
	"strings"
)

func VowelsPresent(line string) bool {

	count_vowels := 0
	for _, vowel := range Vowels {
		if strings.Contains(line, vowel) {
			times_vowel := strings.Count(line, vowel)
			count_vowels += times_vowel
		}
	}

	if count_vowels >= 3 {
		return true
	} else {
		return false
	}
}
