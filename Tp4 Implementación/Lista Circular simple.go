package main

import "fmt"

type Nodo struct {
	value int
	next  *Nodo
}

type LCS struct {
	ref  *Nodo
	size int
}
func ImprimirLista(c *LCS) {
	if c.size == 0 {
		fmt.Println("La lista está vacía.")
		return
	}

	current := c.ref
	for {
		fmt.Printf("%d ", current.value)
		current = current.next
		if current == c.ref {
			break
		}
	}
	fmt.Println()
}
//Igualdad observacional
func (c1 *LCS) Equals(c2 *LCS) bool {
	if c1.EsVacia() && c2.EsVacia() {
		return false
	}
	if c1.size != c2.size {
		return false
	}
	currentc1 := c1.ref
	currentc2 := c2.ref
	for i := 0; i < c1.size; i++ {
		if currentc1.value != currentc2.value {
			return false
		}
		currentc1 = currentc1.next
		currentc2 = currentc2.next
	}
	return true
}

//Generadores

func vacia() *LCS {
	return &LCS{}
}

//Observadores Básicos

func (c *LCS) EsVacia() bool {
	if c.size == 0 {
		return true
	}
	return false
}

func (c *LCS) Tamaño() int {
	return c.size
}

func (c *LCS) Referencia() int {
	if !c.EsVacia(){
		return c.ref.value
	}
	return -1
}

func (c *LCS) Pertenece(a int) bool {
	current := c.ref
	for i := 0; i < c.size; i++ {
		if current.value == a {
			return true
		}
		current = current.next
	}
	return false
}
func (c *LCS) Agregar(a int) {
	nuevoNodo := &Nodo{value: a}
	if c.EsVacia() {
		c.ref = nuevoNodo
		nuevoNodo.next = c.ref // Hacer que el nuevo nodo apunte a sí mismo en una lista circular
	} else {
		nuevoNodo.next = c.ref.next
		c.ref.next = nuevoNodo
	}
	c.size++
}
func(c *LCS)Eliminar(elemento int) bool{
	if c.size==0{
		print("No hay elementos para eliminar")
		return false 
	}
	prev:=c.ref
	for prev.next.value!=elemento && prev.next!=c.ref{ //frena cuando encuentra el siguiente elemento o cuando 
		prev=prev.next 									//llega a referencia (si o sí llega)
	}
	if prev.next == c.ref{ //si el que le sigue es el referencia puede ser que tenga el elemento o que no lo haya encontrado
		if c.ref.value==elemento{ //si el nodo que tiene e	l elemento es el referencia
			if c.size==1{
				print("La referencia eliminada es el único elemento")
				c.ref=nil
				c.size--
				return true
				}
			c.ref=c.ref.next 
			prev.next=prev.next.next
			print("El elemento eliminado fue la referencia")
			c.size--
			return true 
		}else{ //si recorrió toda la cadena y la referencia no tiene el elemento, significa uqe no se encuentra
			print("El elemento no se encuentra en la lista")
			return false
		}
	}else{ //si frenó el for y el siguiente no es el referencia es porque hay otro para eliminar
		prev.next=prev.next.next
		c.size--
		print("El elemento eliminado fue otro")
		return true
	}
}




func main() {
	// Crear una lista circular simple vacía
	LCS := vacia()

	// Verificar si la lista está vacía
	fmt.Println("¿La lista está vacía?", LCS.EsVacia())

	// Agregar elementos a la lista
	LCS.Agregar(10)
	LCS.Agregar(20)
	LCS.Agregar(30)

	// Verificar el tamaño de la lista
	fmt.Println("Tamaño de la lista:", LCS.Tamaño())

	// Verificar si un elemento pertenece a la lista
	fmt.Println("¿El elemento 20 pertenece a la lista?", LCS.Pertenece(20))

	// Mostrar el valor de la referencia
	fmt.Println("Valor de la referencia:", LCS.Referencia())

	// Eliminar un elemento de la lista
	fmt.Println("¿Se eliminó el elemento?", LCS.Eliminar(10))

}
