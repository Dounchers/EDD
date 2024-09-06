// Suma de enteros: permite sumar todos los números enteros comprendidos entre un
// parámetro de inicio y uno de fin.
package main

import (
	"fmt"
)

func SumaEnteros(inicio, fin int) int {
	if inicio == fin {
		return inicio
	} else {
		return fin + SumaEnteros(inicio, fin-1)
	}
}
func main() {
	fmt.Print(SumaEnteros(10, 20))
}
