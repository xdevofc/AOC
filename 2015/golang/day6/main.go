package main

import (
	"bufio"
	"bytes"
	"fmt"
)

var Matrix [1000][1000]bool

func main() {

	// declaramos el nombre del archivo y extraemos los datos
	fileName := "input.txt"
	data, error := OpenFiles(fileName)

	if error != nil {
		panic(error)
	}

	// convertimos []bytes en reader

	reader := bytes.NewReader(data)
	// convertimos el contenido en texto
	scann := bufio.NewScanner(reader)

	// empezamos a recorrer linea a lina

	for scann.Scan() {
		linea := scann.Text()

		// separar instrucciones y coordenadas
		instrucciones, x1, y1, x2, y2 := SplitInfo(linea)

		//realizar accion
		SetValues(instrucciones, x1, y1, x2, y2)
		fmt.Printf("===> %v \n", linea)
		fmt.Printf("%v: (%v,%v) hasta (%v,%v) \n", instrucciones, x1, y1, x2, y2)
	}
	// contar la luces prendidas

	ligthsOn := CountLights()

	// 568658 -> Too low
	// 569004 -> Too low
	// 569999 ->
	fmt.Printf("Luces Prendidas: %v \n", ligthsOn)

}
