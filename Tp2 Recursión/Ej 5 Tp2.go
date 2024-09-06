//Posición mayúscula: dada una cadena de caracteres, devuelve la posición en la que se
//encuentra la primera letra mayúscula.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FirstMayus(c string)(int){
	//base case
	if len(c)==0{ //si la cadena está vacia
		return -1
	}
	if strings.ToLower(c) == c { //si la cadena está en minúsculas
        return -1
    }
	if unicode.IsLetter(rune(c[0])){
		if c[:1]== strings.ToUpper(c[:1]){
			return 0
		}
	}
	return 1 + FirstMayus(c[1:])
}

func FirstMayusIterative(c string)int{
	for i,char := range c{
		if(unicode.IsUpper(char)){
			return i
		}
	}
	return -1
}



func main(){
	fmt.Print(FirstMayus("h oLa Que Tal"))
	fmt.Print(FirstMayusIterative("asfsadfaf"))
}
