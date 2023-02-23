package main

import (
	"S1EDD-C/Clase2/Estructuras"
	"fmt"
)

func main() {
	fmt.Println("X")
	listaDoble := Estructuras.NewLista()
	listaDoble.InsertarAlFinal("auxiliar1", 1)
	listaDoble.InsertarAlFinal("auxiliar2", 2)
	listaDoble.InsertarAlFinal("auxiliar3", 3)
	listaDoble.InsertarAlFinal("auxiliar4", 4)
	listaDoble.InsertarAlFinal("auxiliar5", 5)
	listaDoble.InsertarAlFinal("auxiliar6", 6)
	listaDoble.MostrarConsola()
}
