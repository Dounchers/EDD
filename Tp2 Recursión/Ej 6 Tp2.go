//6. Mayor elemento: dado un arreglo de enteros, devolver el mayor elemento.
package main
import (
	"fmt"
)

func Max(array[]int,index int)int{
	if index == len(array)-1{
		return array[index]
	}
	return Max(array[i],Max(array[i],index+1))

}
func main{
	fmt.Print(BiggestNumber([4,6,8,15],0))
}