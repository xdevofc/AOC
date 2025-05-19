package _2016.java.day1;

import java.util.HashMap;
import java.util.Map;


public class GetInstructions{ 
    

    // variable para guardar las coordenadas  
    Map<String,Integer> Coordenadas = new HashMap<>();

    // variable para guardar la linea
    private String line;
    
    public GetInstructions(String line){
        this.line = line;
    }

    public void ImprimirLinea(){
        System.out.println(line);
    }

    public int GetCoordenates(){
        Coordenadas.put("X", 0);
        Coordenadas.put("Y", 0);
        // separar en un arreglo las instrucciones 
        String[] instrucciones = line.split(", ");
       
        // declarar sentido
        String cabeza;

        if (String.valueOf(instrucciones[0]) == "R"){
            cabeza = "E";
        }else if (String.valueOf(instrucciones[0]) == "L"){
            cabeza = "O";
        }

        return 1;
        
    }
}
