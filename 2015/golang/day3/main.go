package main 

import (
    "fmt"
    "os"
)
// helpers 
func positionExist(c map[string]bool, clave string, posiciones *Coordenadas){
    
    if c[clave] == true{
        fmt.Println("Ya existe")
    }else{
        ContarCasas(posiciones)
        c[clave] = true
    }
}



func main (){
   
    // declaracion de la estructura

    posiciones := Coordenadas{x: 1, y:1, houses:1}
    
    // creamos un map para guardar las coordenadas ingresadas 
    coordenadas_registradas := make(map[string]bool)
    // sprintf formatea la cadena a texto
    clave := fmt.Sprintf("%d,%d", posiciones.x, posiciones.y)

    coordenadas_registradas[clave] = true


    // manipulamos el archivo
    const inputFile = "input.txt"
    
    // leyendo el archivo
    content, err := os.ReadFile(inputFile)
   

    if err != nil {
        panic(err) 
    }

    text := string(content)


    for _, c := range text{
         

        if string(c) == "^" {
            MoverArriba(&posiciones)
            clave := fmt.Sprintf("%d,%d", posiciones.x, posiciones.y)
            positionExist(coordenadas_registradas, clave, &posiciones)
        }else if string(c) == "v"{
            MoverAbajo(&posiciones)
            clave := fmt.Sprintf("%d,%d", posiciones.x, posiciones.y) 
            positionExist(coordenadas_registradas, clave, &posiciones)
        }else if string(c) == ">"{
            MoverDerecha(&posiciones)
            clave := fmt.Sprintf("%d,%d", posiciones.x, posiciones.y)
            positionExist(coordenadas_registradas, clave, &posiciones)
        }else if string(c) == "<"{
            MoverIzquierda(&posiciones)
            clave := fmt.Sprintf("%d,%d", posiciones.x, posiciones.y)
            positionExist(coordenadas_registradas, clave, &posiciones)
        }

             
        fmt.Printf("La posicione en x: %d Y:%d Casas:%d \n",
            posiciones.x,
            posiciones.y,
            posiciones.houses)


// testing

    }

    
    //fmt.Println("TESTING world day3")
    //SecondChallenge()
}
