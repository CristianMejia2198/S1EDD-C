package clase

import (
	"fmt"
	"strconv"
)

type ListaNumero struct {
	Inicio   *nodoLista
	Longitud int
}

func (l *ListaNumero) AgregarNumero(numero []int) {
	for i := 0; i < len(numero); i++ {
		if l.Longitud == 0 {
			nuevoNodo := &nodoLista{valor: numero[i], siguiente: nil, pila: &Pila{primero: nil, longitud: 0}}
			l.Inicio = nuevoNodo
			l.Longitud++
		} else {
			nuevoNodo := &nodoLista{valor: numero[i], siguiente: nil, pila: &Pila{primero: nil, longitud: 0}}
			aux := l.Inicio
			for aux.siguiente != nil {
				aux = aux.siguiente
			}
			aux.siguiente = nuevoNodo
			l.Longitud++
		}
	}
}

func (l *ListaNumero) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println("Posicion: ", aux.valor, " Pila: ", valoresPila(aux.pila))
		aux = aux.siguiente
	}
}

func (l *ListaNumero) AgregarPila(posicion int, cantidad int) {
	if l.Longitud == 0 {
		fmt.Println("No hay elementos")
	} else {
		aux := l.Inicio
		for i := 0; i < l.Longitud; i++ {
			if posicion == aux.valor {
				aux.pila.AgregarValores(cantidad)
				return
			}
			aux = aux.siguiente
		}
		fmt.Println("No se encontro el elemento")
	}
}

func valoresPila(pila *Pila) string {
	contenido := ""
	aux := pila.primero
	if aux != nil {
		for aux != nil {
			contenido += strconv.Itoa(aux.valor)
			contenido += "|"
			aux = aux.siguiente
		}
	} else {
		contenido = "| |"
	}
	return contenido
}
