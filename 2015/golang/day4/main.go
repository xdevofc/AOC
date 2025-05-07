package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main(){

	input := "bgvyzdsv"
	var numero = 1
	
	// creando un loop para encontrar el numero que hace que empieze con 00000 

	for {

		// pasar del numero a string
		num_string := strconv.Itoa(numero)
		
		// unir cadenas
		clave := input + num_string

		// obtener hash directamente
		// encryptando a md5 la cadena de texto
		hash := md5.Sum([]byte(clave))

		// convertir el arreglo a string

		decrypt := hex.EncodeToString(hash[:])


		if strings.HasPrefix(decrypt, "000000"){
			
			fmt.Printf("La clave es: %v\n ", decrypt)
			fmt.Printf("El numero es: %v\n ", numero)
			break
		}

		numero +=1


} 
    
	// SecondChallenge()
}

