// Cuenta regresiva: recibe un número entero e imprime todos los números comprendidos entre el mismo y 0.
package main
import (
	"fmt"
)

func Regresiva(num int) int {
	if num == 0 {
		fmt.Print(0)
		return 0
	} else {
		fmt.Print(num,"\n")
		return Regresiva(num-1)
	}
}

func main(){
	Regresiva(10)
}