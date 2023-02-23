package main

import (
	"S1EDD-C/Clase4/estructuras"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Clase 4")
	//var pila *estructuras.Pila
	pila := &estructuras.Pila{Primero: nil, Longitud: 0}
	//var cola *estructuras.Cola
	cola := &estructuras.Cola{Primero: nil, Longitud: 0}
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("")
		fmt.Println("1. Pila")
		fmt.Println("2. Cola")
		fmt.Println("3. Salir")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			menu_pila(pila)
		case 2:
			menu_cola(cola)
		case 3:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		}
	}

}

func menu_pila(pila *estructuras.Pila) {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("")
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Peek")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Ingresar a la Pila")
			pila.Push(formato_hora())
		case 2:
			fmt.Println("Sacar de la Pila")
			pila.Pop()
		case 3:
			fmt.Println("Mostrar Cima y Graficar")
			pila.Peek()
			pila.Graficar()
		case 4:
			salir = true
		}
	}
}

func menu_cola(cola *estructuras.Cola) {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("")
		fmt.Println("1. Encolar")
		fmt.Println("2. Descolar")
		fmt.Println("3. Primero y Graficar")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Ingresar a la Cola")
			cola.Encolar(formato_hora())
		case 2:
			fmt.Println("Sacar de la Cola")
			cola.Descolar()
		case 3:
			fmt.Println("Mostrar Primero")
			cola.MostrarPrimero()
			cola.Graficar()
		case 4:
			salir = true
		}
	}
}

func formato_hora() string {
	tiempo := time.Now() // 10:04
	texto_final := ""
	if tiempo.Hour() < 10 {
		texto_final = texto_final + "0" + strconv.Itoa(tiempo.Hour()) + ":"
	} else {
		texto_final = texto_final + strconv.Itoa(tiempo.Hour()) + ":"
	}
	if tiempo.Minute() < 10 {
		texto_final = texto_final + "0" + strconv.Itoa(tiempo.Minute()) + ":"
	} else {
		texto_final = texto_final + strconv.Itoa(tiempo.Minute()) + ":"
	}
	if tiempo.Second() < 10 {
		texto_final = texto_final + "0" + strconv.Itoa(tiempo.Second())
	} else {
		texto_final = texto_final + strconv.Itoa(tiempo.Second())
	}
	return texto_final
}
