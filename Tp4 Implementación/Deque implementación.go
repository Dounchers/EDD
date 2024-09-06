package main
import(
	"fmt"
)

type Nodo struct {
	value int
	next  *Nodo
	prev  *Nodo
}
type Deque struct {
	head *Nodo
	tail *Nodo
	size int
}

//Igualdad observacional
func (d1 *Deque) equals(d2 *Deque) bool {
	if d1.isEmpty() && d2.isEmpty() {
		return false
	}
	if d1.size != d2.size {
		return false
	}
	currentd1 := d1.head
	currentd2 := d2.head
	for currentd1 != nil {
		if currentd1.value != currentd2.value {
			return false
		}
		currentd1 = currentd1.next
		currentd2 = currentd2.next
	}
	return true
}

//Generadores
func vacia() *Deque {
	return &Deque{}
}

//Observadores básicos
func (d *Deque) isEmpty() bool {
	if d.size == 0 {
		return true
	}
	return false
}

func (d *Deque) peekSize() int {
	return d.size
}

func (d *Deque) front() int {
	return d.head.value
}
func (d *Deque) back() int {
	return d.tail.value
}

func (d *Deque) pushFront(a int) {
	nuevoNodo := &Nodo{value: a}
	if d.isEmpty() {
		d.head = nuevoNodo
		d.tail = nuevoNodo
		d.size++
	} else {
		d.head.prev = nuevoNodo
		nuevoNodo.next = d.head
		d.head = nuevoNodo
		d.size++
	}
}

func (d *Deque) pushBack(a int) {
	nuevoNodo := &Nodo{value: a}
	if d.isEmpty() {
		d.head = nuevoNodo
		d.tail = nuevoNodo
		d.size++
	} else {
		nuevoNodo.prev = d.tail
		d.tail.next = nuevoNodo
		d.tail = nuevoNodo

	}
}

func (d *Deque) popFront() int {
	if d.isEmpty() {
		// Manejar caso de cola vacía
		println("Error: la cola está vacía")
		return -1 // O algún otro valor de error
	}
	value := d.head.value //en cualquier caso voy a retornar el valor, así que lo guardo y luego opero
	if d.size == 1 {      //si hay un sólo elemento, mi cola estará vacia
		d.head = nil
		d.tail = nil     //desvinculo mi primer elemento dado que si es primero, su sucesor debe ser nil
	} else { 
		d.head = d.head.next   //mi puntero al primero ahora es el siguiente del viejo primero
		d.head.prev.next = nil //mi viejo primero tiene como next nil para que se elimine y desvincule de la cola
		d.head.prev = nil 
	}
	d.size--
	return value
}
func (d *Deque) popBack()int{
	if d.isEmpty() {
		// Manejar caso de cola vacía
		println("Error: la cola está vacía")
		return -1 // O algún otro valor de error
	}
	value := d.head.value //en cualquier caso voy a retornar el valor, así que lo guardo y luego opero
	if d.size == 1 {      //si hay un sólo elemento, mi cola estará vacia
		d.head = nil
		d.tail = nil     
	} else { 
		d.tail = d.tail.prev   // Ahora d.tail apunta al nodo anterior al último nodo
		d.tail.next.prev = nil // El antiguo último sigue apuntando al nuevo último, por lo que se debe cortar la referencia
		d.tail.next = nil      // El siguiente del nuevo último debe apuntar a nil
	}
	d.size--
	return value
}

func main() {
	// Crear una nueva deque
	deque := Deque{}

	// Agregar elementos al frente de la deque
	deque.pushFront(3)
	deque.pushFront(2)
	deque.pushFront(1)

	// Mostrar los elementos de la deque después de agregar al frente
	fmt.Println("Deque después de agregar al frente:")
	printDeque(&deque)

	// Eliminar el primer elemento de la deque
	poppedFront := deque.popFront()
	fmt.Println("\n","Elemento eliminado del frente:", poppedFront)
	fmt.Println("Deque después de eliminar del frente:")
	printDeque(&deque)

	// Agregar elementos al final de la deque
	deque.pushBack(4)
	deque.pushBack(5)

	// Mostrar los elementos de la deque después de agregar al final
	fmt.Println("\n","Deque después de agregar al final:")
	printDeque(&deque)

	// Eliminar el último elemento de la deque
	poppedBack := deque.popBack()
	fmt.Println("\nElemento eliminado del final:", poppedBack)
	fmt.Println("Deque después de eliminar del final:")
	printDeque(&deque)
}

// Función auxiliar para imprimir los elementos de la deque
func printDeque(d *Deque) {
	current := d.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}
