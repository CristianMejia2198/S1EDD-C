package Estructura

import (
	"fmt"
)

/*
Creamos la Estructura Alumno
este tendra 2 atributos nombre y carnet
*/
type Alumno struct {
	nombre string
	carnet int
}

/*
Creamos la estructura del nodo
este tendra 2 atributos, el alumno que hace referencia al nuevo estudiante de Tipo Alumno (estructura creada arriba)
y siguiente que es referencia al siguiente nodo de Tipo Nodo
*/
type Nodo struct {
	alumno    *Alumno
	siguiente *Nodo
}

/*
Creamos la estructura de la lista simple
consta de 2 atributos, inicio que se refiere al primer elemento de la lista de tipo Nodo
y longitud de tipo entero que se refiere al tamaño de la lista actual
*/
type Lista struct {
	inicio   *Nodo
	longitud int
}

/* creamos una Funcion para crear un nuevo nodo de la lista de tipo Nodo
este recibe un alumno de tipo Alumno y retorna el nuevo nodo*/
/* en Go nil hace referencia a un valor null */
func NewNodo(alumno *Alumno) *Nodo {
	return &Nodo{alumno, nil}
}

/* Creamos la Funcion EstaVacia
retornara el valor verdadero si lista esta Vacia (sin elementos)
retornara el valor falso si la lista tiene elementos
se especifica (l *Lista) para hacer referencia que la funcion
pertenece al Struct Lista*/

func (l *Lista) EstaVacia() bool {
	if l.longitud == 0 {
		return true
	} else {
		return false
	}
}

/*
Creamos la funcion AgregarAlumno que recibe de parametro nombre de tipo cadena y carnet de tipo entero
se especifica (l *Lista) para hacer referencia que la funcion
pertenece al Struct Lista
*/
func (l *Lista) AgregarAlumno(nombre string, carnet int) {
	comprobacion := l.EstaVacia() //Corroboramos si esta vacia la lista
	if comprobacion {             //Si esta vacia
		nuevoAlumno := &Alumno{nombre, carnet} // se crea nuevo nodo
		l.inicio = NewNodo(nuevoAlumno)        // se le asigna al inicio de la lista el nuevo nodo
		l.longitud++                           // aumentamos la longitud
	} else { //si ya tiene elementos
		aux := l.inicio // Creamos un auxiliar de la lista
		for aux.siguiente != nil {
			aux = aux.siguiente // en Go no existe el while, se utiliza el for y nos dirigimos al ultimo nodo de la lista
		}
		nuevoAlumno := &Alumno{nombre, carnet}
		aux.siguiente = NewNodo(nuevoAlumno) //se le asigna al siguinte del nodo actual el nuevo nodo
		aux.siguiente.siguiente = nil        //se espeficia que el siguiente del nuevo nodo creado sera valor nulo
		l.longitud++                         //Aumentamos la longitud
	}
}

/**Creamos la funcion MostrarLista para imprimir los valores de la lista*/
func MostrarLista(lista *Lista) {
	auxiliar := lista.inicio
	for auxiliar != nil {
		fmt.Println(auxiliar.alumno.nombre)
		fmt.Println(auxiliar.alumno.carnet)
		fmt.Println("------------------------")
		auxiliar = auxiliar.siguiente
	}
}

/*
Creamos la funcion New_Lista que retorna un objeto de tipo lista
con esto inicializamos la lista y sus valores en el main.go
*/
func New_Lista() *Lista {
	return &Lista{nil, 0}
}
