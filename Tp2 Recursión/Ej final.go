//Dado un arreglo de enteros indicar si está ordenado de forma ascendente.
package main
import "fmt"

func Ordenado(arr[]int)bool{
	
	if len(arr)==0{
		return true
	}
	for j:=0;j<len(arr)-1;j++{
		if arr[j]>arr[j+1]{
			return false
		}
	}
	return Ordenado(arr[1:])
}

func main() {
    enteros := []int{11, 12, 23, 25, 34, 10, 90}
	if(Ordenado(enteros)){
		fmt.Print("El arreglo se encuentra ordenado")
	}else{
		fmt.Print("El arreglo no está ordenado")
	}
}