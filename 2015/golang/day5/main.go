package main

import (
	"bufio"
	"fmt"
	"os"
	"bytes"
)


func main(){
	
	// leyendo un archivo
	fileName := "input.txt"

	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// pasar de []bytes a tipo de io.reader para que pueda leer el scanner
	reader := bytes.NewReader(data)

	// devuelve un scanner del que se tiene que leer
	file_scanner := bufio.NewScanner(reader)

	//imprimir linea a linea
	// TODO [-] FUNCTION TO MATCH TE REQUIREMENTS


	for file_scanner.Scan(){
		linea := file_scanner.Text()
		fmt.Println(linea)
	}

	// en el caso de que no se pueda leer
	if err := file_scanner.Err(); err != nil {
		fmt.Println("Error al leer: ", err)
	}




	fmt.Println("%v", data)


	//content := string(data)



	//fmt.Println("%v",content)
	fmt.Println("%v",Vowels)
}
