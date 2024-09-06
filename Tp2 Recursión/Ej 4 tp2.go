// 4. Vocales y consonantes: devuelve la cantidad de vocales y de consonantes que contiene
// una cadena.
package main

import(
	"fmt"
	"strings"
)
const CONSONANTS="bcdfghjklmnñpqrstvwxyz"
const VOCALS="aeiou"

func VocalsConsonants(c string,totalVocals int, totalConsonants int) (int, int) {

	if len(c) == 0 {
		return totalConsonants,totalVocals
	}
	if strings.ContainsAny(CONSONANTS, c[:1]){
		totalConsonants++
	}
	if strings.ContainsAny(VOCALS,c[:1]){
		totalVocals++
	}
	return VocalsConsonants(c[1:],totalVocals,totalConsonants)
}

func VocalsConsonantsIterative(c string)(int,int){
vocalsIterative:=0
consonantsIterative:=0
	for _,char:=range c{
		if strings.ContainsRune(VOCALS,char){
			vocalsIterative++
		}
		if strings.ContainsRune(CONSONANTS,char){
			consonantsIterative++;
		}	
	}
	return consonantsIterative,vocalsIterative
}
func main(){
	fmt.Print(VocalsConsonants("hola mundo",0,0)) 
	fmt.Print("\n")
	fmt.Print(VocalsConsonantsIterative("hola mundo"))
}

