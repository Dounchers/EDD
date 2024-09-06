//Invertir: dado un arreglo de enteros, invertirlo
package main
import("fmt")
func Invest(arr[]int)[]int{
	if len(arr)<=1{
		return arr
	}
	return append(Invest(arr[1:]), arr[0])
}

func InvestIterative(arr[]int)[]int{
	for i:=0;i<=len(arr)/2;i++{
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func main(){
	arreglo := []int{1, 2, 3, 4, 5}
    fmt.Println("Arreglo original:", arreglo)
    fmt.Println("Arreglo invertido:", Invest(arreglo))
	fmt.Println("Arreglo invertido 2:", InvestIterative((arreglo)))
}

