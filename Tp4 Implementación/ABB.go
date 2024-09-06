package main

import "fmt"

// define la estructura para el Arbol binario de búsqueda
type ArbolBB struct {
	valor              int
	hijo_izq, hijo_der *ArbolBB
}

// crea un arbol binario vacío
func newAbb() *ArbolBB {
	arbol := new(ArbolBB)
	arbol.valor = 0
	arbol.hijo_izq = nil
	arbol.hijo_der = nil
	return arbol
}

func a_partir_de(datos []int) *ArbolBB {
	var arbol *ArbolBB
	arbol = newAbb()
	fin := len(datos)
	for i := 0; i < fin; i = i + 1 {
		arbol.insertar(datos[i])
	}
	return arbol
}

// verifica si el arbol está vacío
func (a *ArbolBB) es_vacio() bool {
	return (a.valor == 0) && (a.hijo_izq == nil) && (a.hijo_der == nil)
}

// inserta un elemento en el árbol
func (a *ArbolBB) insertar(elem int) *ArbolBB {
	if a.es_vacio() {
		a.valor = elem
		a.hijo_izq = new(ArbolBB)
		a.hijo_der = new(ArbolBB)
	} else if elem < a.valor {
		a.hijo_izq.insertar(elem)
	} else if elem > a.valor {
		a.hijo_der.insertar(elem)
	}
	return a
}

func (a *ArbolBB) es_hoja() bool {
	return a.hijo_izq.es_vacio() && a.hijo_der.es_vacio()
}

func (a *ArbolBB) maximo() int {
	if !a.hijo_der.es_vacio() {
		return a.hijo_der.maximo()
	}
	return a.valor
}

func (a *ArbolBB) minimo() int {
	if !a.hijo_izq.es_vacio() {
		return a.hijo_izq.minimo()
	}
	return a.valor
}

func (a *ArbolBB) eliminar(elem int) *ArbolBB {
	if a.es_vacio() {
		fmt.Println("El elemento a eliminar no está en el árbol")
		return a
	} else if elem < a.valor {
		a.hijo_izq.eliminar(elem)
	} else if elem > a.valor {
		a.hijo_der.eliminar(elem)
	} else if (a.valor == elem) && a.es_hoja() {
		a.valor = 0
		a.hijo_izq = nil
		a.hijo_der = nil
		fmt.Println("soy vacio: ", a.es_vacio())
	} else if a.hijo_izq.es_vacio() {
		a.valor = a.hijo_der.valor
		a.hijo_izq = a.hijo_der.hijo_izq
		a.hijo_der = a.hijo_der.hijo_der
	} else if a.hijo_der.es_vacio() {
		a.valor = a.hijo_izq.valor
		a.hijo_izq = a.hijo_izq.hijo_izq
		a.hijo_der = a.hijo_izq.hijo_der
	} else {
		minimo_subArbolDer := a.hijo_der.minimo()
		a.hijo_der.eliminar(minimo_subArbolDer)
		a.valor = minimo_subArbolDer
	}
	return a
}

func (a *ArbolBB) cantidad_de_nodos() int {
	if a.es_vacio() {
		return 0
	} else if a.es_hoja() {
		return 1
	} else {
		return 1 + a.hijo_izq.cantidad_de_nodos() + a.hijo_der.cantidad_de_nodos()
	}
}

// recorridos del árbol
func (a *ArbolBB) preOrder() {
	if !a.es_vacio() {
		fmt.Println(a.valor, " ")
		a.hijo_izq.preOrder()
		a.hijo_der.preOrder()
	}
}

func (a *ArbolBB) inOrder() {
	if !a.es_vacio() {
		a.hijo_izq.inOrder()
		fmt.Println(a.valor, " ")
		a.hijo_der.inOrder()
	}
}


func (a *ArbolBB) postOrder() {
	if !a.es_vacio() {
		a.hijo_izq.postOrder()
		a.hijo_der.postOrder()
		fmt.Println(a.valor, " ")
	}
}

func main() {
	var miArbol, otroArbol *ArbolBB
	miArbol = newAbb()
	miArbol.insertar(13)
	miArbol.insertar(7)
	miArbol.insertar(9)
	miArbol.insertar(31)
	miArbol.insertar(22)
	miArbol.insertar(44)
	fmt.Println("Maximo", miArbol.maximo())
	fmt.Println("Minimo", miArbol.minimo())
	fmt.Println("Cantidad de nodos", miArbol.cantidad_de_nodos())
	secuencia := []int{2, 3, 5, 7, 11, 13}
	otroArbol = a_partir_de(secuencia)
	otroArbol.inOrder()

}