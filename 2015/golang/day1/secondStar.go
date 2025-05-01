package main

import (
	"fmt"
	"os"
)

func ReadInput2() {

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
			fmt.Printf("index: %d\n", i)
            
        }
        
        // calculando en que piso esta
        // en el loop de arriba, la i almacena el index, empezando por el 0
        // por eso al retornar el valor le +1


       pasos_final = pasos_izq - pasos_der
       
       if pasos_final == -1 {
            fmt.Printf("La posicion en la que va al subsuelo es: %d\n", i+1)
            return
       }
	}
	
	//fmt.Printf("Pasos hacia la izq: %d\n", pasos_izq)
	//fmt.Printf("Pasos hacia la der: %d\n", pasos_der)
	//fmt.Printf("Esta en el piso: %d\n", pasos_final)

}
