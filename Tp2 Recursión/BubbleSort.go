package main
import("fmt")

func BubbleSort(arr[]int)[]int{
flag:=false
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
				flag=true
			}	
			
		}
		if(!flag){
			fmt.Print("arreglo ya ordenado")
			return arr
	}
    }
    return arr
}


func BubbleSortRecursive(arr []int, n int) []int {
    flag := false
    for j := 0; j < n-1; j++ {
        if arr[j] > arr[j+1] {
            arr[j], arr[j+1] = arr[j+1], arr[j]
            flag = true
        }
    }
    if !flag {
        fmt.Println("Arreglo ya ordenado")
        return arr
    }
    return BubbleSortRecursive(arr, n-1)
}
func main(){
	arreglo := []int{11, 12, 22, 25, 34, 64, 90}
    fmt.Println("Arreglo original:", arreglo)
    fmt.Println("Arreglo ordenado: ",BubbleSort(arreglo))
	fmt.Println("Arreglo ordenado 2:", BubbleSortRecursive(arreglo,len(arreglo)))
}

