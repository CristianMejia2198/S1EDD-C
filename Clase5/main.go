package main

import (
	"Clase5/clase"
	"fmt"
)

func main() {
	var ruta string
	fmt.Println("Clase 5")
	fmt.Println("")
	fmt.Println("Parte 1 - Lectura CSV")
	fmt.Scanln(&ruta)
	clase.LeerArchivo(ruta)
	fmt.Println("")
	fmt.Println("Parte 2 - Ordenamiento de Burbuja")
	numeros := [10]int{4, 6, 5, 3, 1, 9, 15, 10, 89, 78}
	fmt.Println("Lista sin ordenar: ", numeros)
	clase.Ordenamiento_Burbuja_Arreglo(numeros)
	fmt.Println("")
	fmt.Println("Parte 3 - Ordenamiento de burbuja para Lista Dinamica")
	lista := &clase.ListaNumero{Inicio: nil, Longitud: 0}
	lista.AgregarNumero(numeros[:])
	fmt.Println(" LISTA NO ORDENADA ")
	lista.Mostrar()
	fmt.Println(" LISTA YA ORDENADA")
	clase.Ordenamiento_Burbuja(lista)
	lista.AgregarPila(1, 4)
	lista.AgregarPila(4, 3)
	lista.AgregarPila(10, 9)
	lista.Mostrar()
	contenido := clase.ArchivoJSON(lista)
	clase.CrearArchivo()
	clase.EscribirArchivo(contenido)
}
