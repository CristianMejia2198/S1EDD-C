package main

import (
	"S1EDD-C/Clase1/Estructura"
	"fmt"
)

func main() {
	listaNueva := Estructura.New_Lista()            //Creamos nueva lista
	listaNueva.AgregarAlumno("Cristian", 201700000) //Agregamos alumno 1
	listaNueva.AgregarAlumno("Alberto", 201700001)  //Agregamos alumno 2
	Estructura.MostrarLista(listaNueva)             // Visualizamos los elementos de la lista
	fmt.Println("XXXXXXXXXXX")
}
