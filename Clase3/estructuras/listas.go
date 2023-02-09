package estructuras

import "fmt"

// Estructura de una Lista Doblemente Enlazada
type ListaDoble struct {
	Inicio   *Nodo
	Final    *Nodo
	Longitud int
}

// Estructura de una Lista Circular Doblemente Enlazada
type ListaCircularDoble struct {
	Inicio   *Nodo
	Final    *Nodo
	Longitud int
}

// Funciones estaVacia de las listas anteriormente declaradas
// Ambas funciones no pueden ser exportadas (quiere decir son privadas)
func (l *ListaDoble) estaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *ListaCircularDoble) estaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

// Funciones para la creacion de los nuevos nodos para las listas
// Ambas funciones no pueden ser exportadas (quiere decir son privadas)
func (l *ListaDoble) newNodo(empleado *Empleado) *Nodo {
	return &Nodo{
		empleado,
		nil,
		nil,
	}
}

func (l *ListaCircularDoble) newNodo(empleado *Empleado) *Nodo {
	return &Nodo{
		empleado,
		nil,
		nil,
	}
}

// Funcion AgregarEmpleado que se puede exportar (es decir es publica)
// Esta funcion es para la Lista Doblemente Enlazada
func (l *ListaDoble) AgregarEmpleado(nombre string, edad int) {
	nuevoEmpleado := &Empleado{nombre, edad}
	if l.estaVacia() {
		nuevoNodo := l.newNodo(nuevoEmpleado)
		l.Inicio = nuevoNodo
		l.Final = nuevoNodo
		l.Longitud++
	} else {
		nuevoNodo := l.newNodo(nuevoEmpleado)
		if l.Final.anterior == nil { //corroboramos si hay un solo elemento
			nuevoNodo.anterior = l.Inicio
			l.Inicio.siguiente = nuevoNodo
			l.Final = nuevoNodo
		} else { //Si hay mas de 1 elemento
			l.Final.siguiente = nuevoNodo
			nuevoNodo.anterior = l.Final
			l.Final = nuevoNodo
		}
		l.Longitud++
	}
}

// Funcion AgregarEmpleado que se puede exportar (es decir es publica)
// Esta funcion es para la Lista Circular Doblemente Enlazada
func (l *ListaCircularDoble) AgregarEmpleado(nombre string, edad int) {
	nuevoEmpleado := &Empleado{nombre, edad}
	if l.estaVacia() {
		nuevoNodo := l.newNodo(nuevoEmpleado)
		l.Inicio = nuevoNodo
		l.Final = nuevoNodo
		l.Inicio.siguiente = l.Final
		l.Inicio.anterior = l.Final
		l.Final.siguiente = l.Inicio
		l.Final.anterior = l.Inicio
		l.Longitud++
	} else {
		nuevoNodo := l.newNodo(nuevoEmpleado)
		if l.Longitud == 1 {
			nuevoNodo.anterior = l.Inicio // 1 <- 2 ->
			nuevoNodo.siguiente = l.Inicio
			l.Inicio.siguiente = nuevoNodo
			l.Final = nuevoNodo
		} else {
			nuevoNodo.siguiente = l.Inicio
			l.Final.siguiente = nuevoNodo
			nuevoNodo.anterior = l.Final
			l.Final = nuevoNodo
		}
		l.Longitud++
	}
}

// Funciones para mostrar  el contenido de ambas listas
func (l *ListaDoble) MostrarLista() {
	aux := l.Inicio
	for aux != nil {
		fmt.Printf("Nombre: %s, edad: %d \n", aux.empleado.nombre, aux.empleado.edad)
		aux = aux.siguiente
	}
}

func (l *ListaCircularDoble) MostrarLista() {
	aux := l.Final
	for i := 0; i < l.Longitud; i++ {
		fmt.Printf("Nombre: %s, edad: %d \n", aux.empleado.nombre, aux.empleado.edad)
		aux = aux.anterior
	}
}
