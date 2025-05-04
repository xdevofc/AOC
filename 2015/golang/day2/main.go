package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CalcWrap(L int, A int, H int) int {

	// declaro una lista con los elementos a comparar
	lista := []int{L * A, A * H, L * H}

	// ordeno la lista de menor a mayor

	sort.Ints(lista)

	//obtener los primeros elementos
	minor := lista[0]

	res := minor + (2 * L * A) + (2 * A * H) + (2 * H * L)
	//fmt.Printf("RES CALCWRAP: %d\n",res)
	return res
}

func main() {
	// almacenando el papel necesario
	var totalWrap int = 0
	// calcular la cina necesaria
	var totalRibbon int = 0
	// leer el archivo y extraer los datos
	var fileName string = "input.txt"

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close() // Cerrar el archivo

	// creando un archivo para leer
	scanner := bufio.NewScanner(file)

	// leyendo linea por linea
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")

		if len(parts) != 3 {
			fmt.Println("Formato incorrecto: ", line)
			continue
		}

		// convertir cada parte a entero

		L, err1 := strconv.Atoi(parts[0])
		A, err2 := strconv.Atoi(parts[1])
		H, err3 := strconv.Atoi(parts[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Error en conversion", line)
			continue
		}

		//imprimiendo la linea

		fmt.Println("Linea: ", line)

		// calculando cuanto papel por regalo
		res := CalcWrap(L, A, H)
		res2 := SecondChallenge(L, A, H)
		totalWrap = res + totalWrap
		totalRibbon = res2 + totalRibbon
	}

	fmt.Printf("Total Wrap:%d \n", totalWrap)
	fmt.Printf("Total Ribbon:%d \n", totalRibbon)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

}
