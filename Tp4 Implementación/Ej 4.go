package main

import (
	"fmt"
	"strings"
)

type Nodo struct {
	valor    rune
	siguiente *Nodo
}

type Pila struct {
	tope *Nodo
	size int
}
func (s1 *Pila) Equals(s2 *Pila) bool {
	if s1.size != s2.size {
		return false // Si las pilas tienen diferente tamaño, ya no pueden ser iguales
	}
	nodoS1 := s1.tope
	nodoS2 := s2.tope

	for nodoS1 != nil {
		if nodoS1.valor != nodoS2.valor {
			return false // Los elementos de los nodos no son iguales
		}
		nodoS1 = nodoS1.siguiente //avanzo en las pilas
		nodoS2 = nodoS2.siguiente
	}
	return true //si pasa por las anteriores verificaciones, entonces las pilas son iguales
}

func (s *Pila) esVacia() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (p *Pila) Push(valor rune) {
	nuevoNodo := &Nodo{
		valor:    valor,
		siguiente: p.tope,
	}
	p.tope = nuevoNodo
	p.size++
}

func (p *Pila) Pop() rune {
	if p.size == 0 {
		return 0
	}
	valor := p.tope.valor
	p.tope = p.tope.siguiente
	p.size--
	return valor
}

func (p *Pila) Size() int {
	return p.size
}

func main() {
	secuencia:="radar"
	palindromo:=true;

	// Eliminar espacios en blanco y convertir a minúsculas
	secuencia = strings.ToLower(strings.ReplaceAll(secuencia, " ", ""))

	// Crear una pila para almacenar los caracteres de la secuencia
	pila := &Pila{}
	for _, char := range secuencia {
		pila.Push(char)
	}
	current:=pila.tope
	for i:=0;i<len(secuencia);i++{
		if(current.valor != rune(secuencia[i])){
			palindromo=false;
			break
		}
		current=current.siguiente
	}
	if(palindromo){
		fmt.Print("Es palíndromo anashe")
	}else{
		fmt.Print("No es palíndromo unu")
	}

}
