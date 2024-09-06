package main

import "fmt"

type Conjunto struct {
	items []int
}

//GENERADORES
func vacio() *Conjunto {
	return &Conjunto{}
}

func aPartirDe(secuencia []int) *Conjunto {
	nuevoConjunto := vacio()
	for _, elemento := range secuencia {
		nuevoConjunto.Agregar(elemento)
	}
	return nuevoConjunto
}

//IGUALDAD OBSERVACIONAL
func (c1 *Conjunto) equals(c2 *Conjunto) bool {
	if c1.esVacio() != c2.esVacio() {
		return false
	}
	if c1.tamaño() != c2.tamaño() {
		return false
	}
	if !c1.diferencia(c2).esVacio() {
		return false
	}
	return true
}

//OBSERVADORES BÁSICOS
func (c *Conjunto) esVacio() bool {
	if len(c.items) == 0 {
		return true
	}
	return false
}

func (c *Conjunto) Pertenece(elemento int) bool {
	for _, e := range c.items {
		if e == elemento {
			return true
		}
	}
	return false
}
func (c *Conjunto) tamaño() int {
	return len(c.items)
}

//OTRAS OPERACIONES
func (c *Conjunto) Agregar(elemento int) {
	if !c.Pertenece(elemento) { //si no pertenece lo agrego al conjunto
		c.items = append(c.items, elemento)
	}
}
func (c *Conjunto) Eliminar(elemento int) {
	for i, e := range c.items {
		if e == elemento {
			c.items = append(c.items[:i], c.items[i+1:]...)
			break
		}
	}
}
func Union(conjunto1, conjunto2 *Conjunto) *Conjunto {
	union := &Conjunto{}
	// Agregar elementos del primer conjunto
	for _, elemento := range conjunto1.items {
		union.Agregar(elemento)
	}

	// Agregar elementos del segundo conjunto que no estén en la unión
	for _, elemento := range conjunto2.items {
		union.Agregar(elemento)
	}
	return union
}

func (c1 *Conjunto) Interseccion(c2 *Conjunto) *Conjunto {
	interseccion := &Conjunto{}
	for _, elemento := range c1.items {
		if c2.Pertenece(elemento) {
			interseccion.Agregar(elemento)
		}
	}
	return interseccion
}

func (c1 *Conjunto) diferencia(c2 *Conjunto) *Conjunto {
	diferencia := &Conjunto{}
	for _, elemento := range c1.items { //por cada elemento de c1
		if !c2.Pertenece(elemento) { //si el elemento no está en c2, lo agrega a diferencia
			diferencia.Agregar(elemento) //agrega los elementos que estan en el primero pero no estan en el segundo
		}
	}
	return diferencia
}

func main() {
	// Crear conjuntos a partir de secuencias
	conjVacio := vacio()
	print("¿Es vacío?", conjVacio.esVacio())
	secuencia1 := []int{1, 2, 3, 4}
	secuencia2 := []int{3, 4, 5, 6}
	conjunto1 := aPartirDe(secuencia1)
	conjunto2 := aPartirDe(secuencia2)

	// Verificar si los conjuntos son iguales
	fmt.Println("¿Conjuntos iguales?", conjunto1.equals(conjunto2)) // Debería ser false

	// Calcular la unión de los conjuntos
	union := Union(conjunto1, conjunto2)
	fmt.Println("Unión de conjuntos:", union.items) // Debería imprimir [1 2 3 4 5 6]

	// Calcular la intersección de los conjuntos
	interseccion := conjunto1.Interseccion(conjunto2)
	fmt.Println("Intersección de conjuntos:", interseccion.items) // Debería imprimir [3 4]

	// Calcular la diferencia entre los conjuntos
	diferencia := conjunto1.diferencia(conjunto2)
	fmt.Println("Diferencia de conjuntos:", diferencia.items) // Debería imprimir [1 2]

	// Agregar un elemento y verificar su pertenencia
	conjunto1.Agregar(5)
	fmt.Println("¿El elemento 5 pertenece al conjunto 1?", conjunto1.Pertenece(5)) // Debería ser true

	// Eliminar un elemento y verificar su pertenencia
	conjunto2.Eliminar(3)
	fmt.Println("¿El elemento 3 pertenece al conjunto 2?", conjunto2.Pertenece(3)) // Debería ser false
}
