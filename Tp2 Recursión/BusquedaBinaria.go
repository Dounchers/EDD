package main
import(
	"fmt"
	"math"
)

func BusquedaBinaria(numeros []int, valor int) int {
	inicio := 0
	fin := len(numeros)
	for inicio < fin {
	puntoMedio := int(math.Floor(float64(inicio+fin) / 2))
	switch {
	case numeros[puntoMedio] == valor:
	return puntoMedio
	case numeros[puntoMedio] < valor:
	inicio = puntoMedio + 1
	default:
	fin = puntoMedio
	}
	}
	return -1
	}

func BusquedaBinariaRecursiva(numeros []int, valor int, inicio, fin int) int {
		if inicio > fin {
        	return -1
    	}
		puntoMedio := int(math.Floor(float64(inicio+fin) / 2))
		switch {
		case numeros[puntoMedio] == valor:
		return puntoMedio
		case numeros[puntoMedio] < valor:
		inicio = puntoMedio + 1
		default:
		fin = puntoMedio
		}
		return BusquedaBinariaRecursiva(numeros, valor, inicio,fin)
		}

func main(){
	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Arreglo original:", arreglo)
	fmt.Println("Indice del elemento buscado: ", BusquedaBinaria(arreglo,12))
	fmt.Println("Indice del elemento buscado: ", BusquedaBinariaRecursiva(arreglo,10,0,len(arreglo)))
}