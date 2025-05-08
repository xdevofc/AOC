package main

import (
	"strings"
)

func ForbidenMatch(line string) bool {

	if strings.Contains(line, Prohib[0]) || strings.Contains(line, Prohib[1]) ||
		strings.Contains(line, Prohib[2]) || strings.Contains(line, Prohib[3]) {
		return false
	}
	return true
}
