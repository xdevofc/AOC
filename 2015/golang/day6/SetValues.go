package main

import (
	"fmt"
	"strconv"
)

func SetValues(action string, iniciox string, inicioy string, finx string, finy string) {

	// pasando las coordenadas a numeros
	x1, errx1 := strconv.Atoi(iniciox)
	y1, erry1 := strconv.Atoi(inicioy)
	x2, errx2 := strconv.Atoi(finx)
	y2, erry2 := strconv.Atoi(finy)

	// manejando errores
	if errx1 != nil || errx2 != nil || erry1 != nil || erry2 != nil {
		fmt.Println("ERROR EN LA CONVERSION:")
	}

	if action == "turn on" {
		// prendiendo las luces
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				Matrix[i][j] = true
			}
		}
	} else if action == "turn off" {
		// apagando las luces
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				Matrix[i][j] = false
			}
		}

	} else if action == "toggle" {
		// cambiando los valoreslas luces
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				Matrix[i][j] = !Matrix[i][j]
			}

		}
	}
}
