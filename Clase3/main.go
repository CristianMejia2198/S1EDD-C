package main

import (
	"S1EDD-C/Clase3/estructuras"
	"fmt"
)

func main() {
	fmt.Println("Clase 3")

	// Declaramos 2 variables lista1 y lista2, ambos de tipo Operaciones, la cual es la interfaz que usaremos
	// luego a cada lista se inicializa su estructura, NOTA: no se puede inicializar directamente en la estructura
	var lista1 estructuras.Operaciones
	lista1 = &estructuras.ListaDoble{Inicio: nil, Final: nil, Longitud: 0}
	var lista2 estructuras.Operaciones
	lista2 = &estructuras.ListaCircularDoble{Inicio: nil, Final: nil, Longitud: 0}

	// Se crean las variables que utilizaremos en el menu para la insercion
	var (
		nombre string
		edad   int
	)
	opcion := 0    // Nos servira para elegir una opcion en el menu
	salir := false // Para que el menu este en bucle

	for !salir {
		fmt.Println("Menu - Clase 3")
		fmt.Println("1. Ingresar trabajador")
		fmt.Println("2. Ver Lista 1")
		fmt.Println("3. Ver Lista 2")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opcion: ")
		// Scanln es una funcion nativa que nos permite capturar datos desde consola
		//para tomar el dato siempre se coloca & seguido el nombre de variable siempre
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Ingrese un nombre: ")
			fmt.Scanln(&nombre)
			fmt.Println("Ingrese la edad: ")
			fmt.Scanln(&edad)
			if edad < 25 {
				lista1.AgregarEmpleado(nombre, edad)
			} else {
				lista2.AgregarEmpleado(nombre, edad)
			}
		case 2:
			lista1.MostrarLista()
		case 3:
			lista2.MostrarLista()
		case 4:
			fmt.Println("Cerrando Aplicacion....")
			salir = true
		}
	}
}
