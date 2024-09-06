package main

import (
	"fmt"
)

//parámetro formal

type Nodo struct {
	elemento  int
	siguiente *Nodo
}

type Cola struct {
	primero *Nodo
	ultimo  *Nodo
	tamaño  int
}

// Imprimir cola
func mostrarCola(c *Cola) {
	fmt.Print("primero -> ")
	current := c.primero
	for current != nil {
		fmt.Print(current.elemento)
		if current.siguiente != nil {
			fmt.Print(" -> ")
		}
		current = current.siguiente
	}
	fmt.Println(" <- ultimo")
}

//Igualdad observacional

func (c1 *Cola) equals(c2 *Cola) bool {
	if c1.esVacia() == c2.esVacia() {
		print("Ambas listas están vacías")
		return false
	}
	if c1.tamaño != c2.tamaño {
		return false
	}
	nodoC1 := c1.primero
	nodoC2 := c2.primero

	for nodoC1 != nil {
		if nodoC1.elemento != nodoC2.elemento {
			return false
		}
		nodoC1 = nodoC1.siguiente
		nodoC2 = nodoC2.siguiente
	}
	return true
}

// Observadores básicos
func (c *Cola) size() int {
	return c.tamaño
}

func (c *Cola) esVacia() bool {
	return (c.primero == nil && c.ultimo == nil)
}

func (c *Cola) primerElemento() int {
	if c.esVacia() {
		return -1
	}
	return c.primero.elemento
}
func (c *Cola) ultimoElemento() int {
	if c.esVacia() {
		return  -1
	}
	return c.ultimo.elemento
}

// Generadores
func vacia() *Cola {
	return &Cola{} // Devuelve un puntero a una nueva instancia de Cola
}

// Otras operaciones
func (c *Cola) desencolar() int {
	if c.esVacia() {
		return -1
	}
	value := c.primerElemento()
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	c.tamaño--
	return value
}

func (c *Cola) encolar(elemento int) {
	nuevoNodo := &Nodo{elemento: elemento, siguiente: nil} // Crear un nuevo nodo con el elemento dado
	// Si la cola está vacía, el nuevo nodo se convierte en el primer nodo de la cola
	if c.esVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
		c.ultimo = nuevoNodo
	}
	c.tamaño++
}

func main() {
	miCola := vacia()
	miCola2 := vacia()
	print(miCola.equals(miCola2))
	print(miCola.esVacia(), "\n")
	miCola.encolar(5)
	miCola.encolar(6)
	miCola.encolar(7)
	miCola.encolar(9)
	mostrarCola(miCola)
	print("\n")
	print("tamaño de la cola: ", miCola.size(), "\n")
	print("elemento desencolado: ", miCola.desencolar(), "\n")
	mostrarCola(miCola)
	print("elemento desencolado: ", miCola.desencolar(), "\n")
	mostrarCola(miCola)
	print("elemento desencolado: ", miCola.desencolar(), "\n")
	mostrarCola(miCola)
	print("elemento desencolado: ", miCola.desencolar(), "\n")
	mostrarCola(miCola)
	print("elemento desencolado: ", miCola.desencolar(), "\n")
	mostrarCola(miCola)
}
