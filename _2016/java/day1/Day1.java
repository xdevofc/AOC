package _2016.java.day1;

import java.io.File; // importamos la clase de file 
import java.io.FileNotFoundException; // importamos la clase para manejar errores
import java.util.Scanner; // clase para leeer las lineas 

public class Day1{
    public static void main(String[] args){
        // leemos nuevo archivo
        try {
            File myFile = new File("_2016/java/day1/input.txt");
            Scanner myReader = new Scanner(myFile);
            while (myReader.hasNextLine()){
                String data = myReader.nextLine();
                GetInstructions gi = new GetInstructions(data);
                gi.GetCoordenates();
                //System.out.println(data);
            }
            // cerramos el archivo
            myReader.close();
            // manejamos el error
        } catch (FileNotFoundException e) {
            // TODO: handle exception
            System.out.println("Error al leer el archivo");
            e.printStackTrace();
        }
    }
}