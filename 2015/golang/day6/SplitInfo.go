package main

import (
	"regexp"
	"strings"
)

func SplitInfo(line string) (string, string, string, string, string) {
	info := ""
	// extraemos que accion es la que hay que hacer
	if strings.HasPrefix(line, "turn on ") {
		info = "turn on"
	} else if strings.HasPrefix(line, "turn off ") {
		info = "turn off"
	} else if strings.HasPrefix(line, "toggle ") {
		info = "toggle"
	}

	// extrayendo las coordenadas con expresion regular
	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`) // es una expresion regular que extrae las coordenadas

	matches := re.FindStringSubmatch(line) // -> Devuelve en un arreglo pos 0, el input, en el los otros index las coordenadas extraidas

	x1 := matches[1]
	y1 := matches[2]
	x2 := matches[3]
	y2 := matches[4]

	return info, x1, y1, x2, y2
}
