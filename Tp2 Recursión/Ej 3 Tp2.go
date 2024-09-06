//Factorial de un número entero.

package main
import(
	"fmt"
)
func Factorial(num int)int{
	if num==0{
		return 1
	}else{
		return num*Factorial(num-1)
	}
}
func main(){
	fmt.Print(Factorial(5))
}