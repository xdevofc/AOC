package main

import (
	"strings"
)

func Pattern(line string) bool {
	count_pattern := 0
	// empezando a comparar si es que existen patrones de dos letras que se repiten
	for i := 0; i < len(line)-2; i++ {

		finish := i + 2
		start := i
		prefix := line[start:finish]

		count_pattern = strings.Count(line, prefix)

		if count_pattern >= 2 {
			return true
		}

	}
	return false
}
