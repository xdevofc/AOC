package main

import (
	"fmt"
	"os"
)

func ReadInput() {

	// creando los contadores de pasos y resultado
	var pasos_izq int = 0
	var pasos_der int = 0
	var pasos_final int = 0

	// abriendo el txt
	fmt.Println("Reading file")
	fileName := "input.txt"

	//extrayendo los datos
	data, err := os.ReadFile(fileName)

	// handling en caso de error
	if err != nil {
		panic(err)
	}

	// convertir el contenido a string
	text := string(data)

	// Leer los datos del txt
	for i, c := range text {

		if c == '(' {
			pasos_izq++
		} else if c == ')' {
			pasos_der++
		} else {
			fmt.Println("type %s posicion: %d", c, i)
		}

	}

	//calcular el piso
	if pasos_izq > pasos_der {
		pasos_final = pasos_izq - pasos_der
	} else if pasos_der > pasos_izq {
		pasos_final = pasos_der - pasos_izq
	} else {
		pasos_final = 0
	}

	fmt.Printf("Pasos hacia la izq: %d\n", pasos_izq)
	fmt.Printf("Pasos hacia la der: %d\n", pasos_der)
	fmt.Printf("Esta en el piso: %d\n", pasos_final)
}

func main() {
	ReadInput()
}
