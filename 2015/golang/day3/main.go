package main 

import (
    "fmt"
    "os"
)
// helpers 
func positionExist(c map[string]bool, clave string, position *Coordenadas){
    
    if c[clave] == true{
        fmt.Println("Ya existe")
    }else{
        ContarCasas(position)
        c[clave] = true
    }
}



func main (){
   
    // declaracion de la estructura

    posicionesSanta := Coordenadas{x: 1, y:1, houses:0}
    posicionesRobot := Coordenadas{x:1, y:1, houses:0}
    
    // creamos un map para guardar las coordenadas ingresadas 
    
    coordenadas_registradas := make(map[string]bool)
    // sprintf formatea la cadena a texto
    clave := fmt.Sprintf("%d,%d", posicionesSanta.x, posicionesSanta.y)
    
    // creamos el camino que recorre santa y el roboto 
    coordenadas_registradas[clave] = true
    ContarCasas(&posicionesRobot)
    ContarCasas(&posicionesSanta)
    //turno 

    turno := 0


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
        // verificar quien da el siguiente paso
            if turno == 0 {
                MoverArriba(&posicionesSanta)
                clave := fmt.Sprintf("%d,%d", posicionesSanta.x, posicionesSanta.y)
                positionExist(coordenadas_registradas, clave, &posicionesSanta)
                turno +=1
            }else{
                MoverArriba(&posicionesRobot)
                claveRobot := fmt.Sprintf("%d,%d", posicionesRobot.x, posicionesRobot.y)
                positionExist(coordenadas_registradas, claveRobot, &posicionesRobot)
                turno -= 1 
            }
           

        }else if string(c) == "v"{
                 // verificar quien da el siguiente paso
                if turno == 0 {
                    MoverAbajo(&posicionesSanta)
                    clave := fmt.Sprintf("%d,%d", posicionesSanta.x, posicionesSanta.y) 
                    positionExist(coordenadas_registradas, clave, &posicionesSanta)
                    turno +=1 
               }else{
                    MoverAbajo(&posicionesRobot)
                    claveRobot := fmt.Sprintf("%d,%d", posicionesRobot.x, posicionesRobot.y)
                    positionExist(coordenadas_registradas, claveRobot, &posicionesRobot)
                    turno -= 1 
                }


        }else if string(c) == ">"{
                 // verificar quien da el siguiente paso
                if turno == 0 {
                    MoverDerecha(&posicionesSanta)
                    clave := fmt.Sprintf("%d,%d", posicionesSanta.x, posicionesSanta.y) 
                    positionExist(coordenadas_registradas, clave, &posicionesSanta)
                    turno +=1 
               }else{
                   MoverDerecha(&posicionesRobot)
                    claveRobot := fmt.Sprintf("%d,%d", posicionesRobot.x, posicionesRobot.y)
                    positionExist(coordenadas_registradas, claveRobot, &posicionesRobot)
                    turno -= 1 
                }

        }else if string(c) == "<"{
                 // verificar quien da el siguiente paso
                if turno == 0 {
                    MoverIzquierda(&posicionesSanta)
                    clave := fmt.Sprintf("%d,%d", posicionesSanta.x, posicionesSanta.y) 
                    positionExist(coordenadas_registradas, clave, &posicionesSanta)
                    turno +=1 
               }else{
                   MoverIzquierda(&posicionesRobot)
                    claveRobot := fmt.Sprintf("%d,%d", posicionesRobot.x, posicionesRobot.y)
                    positionExist(coordenadas_registradas, claveRobot, &posicionesRobot)
                    turno -= 1 
                }

        }
        

             
        fmt.Printf("La posiciones SANTA: en x: %d Y:%d Casas:%d \n",
            posicionesSanta.x,
            posicionesSanta.y,
            posicionesSanta.houses)

     
            fmt.Printf("La posiciones ROBOT: en x: %d Y:%d Casas:%d \n",
            posicionesRobot.x,
            posicionesRobot.y,
            posicionesRobot.houses)

            
            fmt.Println("La suma de casas es:", len(coordenadas_registradas))


    }

    
    //fmt.Println("TESTING world day3")
    //SecondChallenge()
}
