package main

import ()

func RepeatedLetters(line string) bool {

	// recorrer letra a letra
	for i := 1; i < len(line); i++ {
		if line[i-1] == line[i] {
			return true
		}
	}
	return false
}
