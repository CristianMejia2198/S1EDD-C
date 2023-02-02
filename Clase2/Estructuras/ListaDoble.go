package Estructuras

import "fmt"

type ListaDoble struct {
	inicio *Nodo
	//final *Nodo
	longitud int
}

func (l *ListaDoble) estaVacia() bool {
	if l.longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *ListaDoble) newNodo(alumno *Alumno) *Nodo {
	return &Nodo{alumno, nil, nil}
}

func (l *ListaDoble) InsertarAlFinal(nombre string, id int) {
	nuevoAlumno := &Alumno{nombre, id}
	if l.estaVacia() {
		l.inicio = l.newNodo(nuevoAlumno)
		l.longitud++
	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		//  null <- 1 -> <- 2 -> nil
		aux.siguiente = l.newNodo(nuevoAlumno)
		aux.siguiente.anterior = aux
		l.longitud++
	}
}

func (l *ListaDoble) MostrarConsola() {
	aux := l.inicio
	for aux != nil {
		fmt.Print(aux.alumno.id)
		fmt.Println(" ---> ", aux.alumno.nombre)
		aux = aux.siguiente
	}
}

func NewLista() *ListaDoble {
	lista := new(ListaDoble)
	lista.inicio = nil
	lista.longitud = 0
	return lista
}
