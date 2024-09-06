package main

import (
	"fmt"
)

// imprimir pila
func mostrarPila(s *Pila) {
	current := s.tope
	for current != nil {
		fmt.Println("↓")
		fmt.Println(current.elemento)
		current = current.siguiente // avanzar al siguiente nodo
	}
}

// Parámetro Formal
type Nodo struct { //Creación de la estructura de un nodo
	elemento  int
	siguiente *Nodo
}

type Pila struct { //creación de la pila, de la cual se conoce su tope y su tamaño
	tope   *Nodo
	tamaño int
}

func (s1 *Pila) equals(s2 *Pila) bool {
	if s1.esVacia() && s2.esVacia(){
		return false
	}

	if s1.tamaño != s2.tamaño {
		return false // If stacks have different sizes, they cannot be equal
	}

	nodoS1 := s1.tope
	nodoS2 := s2.tope

	for nodoS1 != nil {
		if nodoS1.elemento != nodoS2.elemento {
			return false // If any pair of elements is different, stacks are not equal
		}
		nodoS1 = nodoS1.siguiente
		nodoS2 = nodoS2.siguiente
	}
	return true // If all elements matched, stacks are equal
}

// Observadores básicos
func (s *Pila) size() int {
	return s.tamaño
}

func (s *Pila) esVacia() bool {
	if s.tamaño == 0 {
		return true
	}
	return false
}

func (s *Pila) topePila() int {
	return s.tope.elemento
}

// Generadores
func vacia() *Pila {
	return &Pila{}
}

func (s *Pila) aPartirDe(secuencia []int) *Pila {
	pila := new(Pila)
	for _, elemento := range secuencia {
		pila.push(elemento)
		pila.tamaño++
	}
	return pila
}

// Otras operaciones
func (s *Pila) push(n int) {
	nuevoNodo := &Nodo{ // Crear un nuevo nodo
		elemento:  n,
		siguiente: s.tope, // El siguiente del nuevo nodo es el actual tope de la pila
	}
	s.tope = nuevoNodo // Actualizar el tope de la pila al nuevo nodo creado
	s.tamaño++         // Incrementar el tamaño de la pila
}

func (s *Pila) pop() int {
	if s.esVacia() {
		return -1
	}
	elementoAux := s.tope.elemento //antes de borrar el nodo guardo el valor que contiene y lo retorno
	s.tope = s.tope.siguiente
	s.tamaño--
	return elementoAux
}

