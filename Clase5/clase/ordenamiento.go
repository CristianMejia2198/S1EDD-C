package clase

import "fmt"

func Ordenamiento_Burbuja(lista *ListaNumero) {
	if lista.Inicio != nil {
		pivote := lista.Inicio
		contador := 0
		for contador != lista.Longitud {
			actual := pivote.siguiente
			for actual != nil {
				if pivote.valor > actual.valor {
					tempValor := pivote.valor
					tempPila := pivote.pila
					pivote.valor = actual.valor
					pivote.pila = actual.pila
					actual.valor = tempValor
					actual.pila = tempPila
				}
				actual = actual.siguiente
			}
			pivote = pivote.siguiente
			contador++
		}
	} else {
		fmt.Println("No existe valores en la lista")
	}
}

func Ordenamiento_Burbuja_Arreglo(listado [10]int) {
	lista := listado

	for i := 0; i < len(lista); i++ {
		for j := 0; j < len(lista)-1; j++ {
			if lista[j] > lista[j+1] {
				temp := lista[j]
				lista[j] = lista[j+1]
				lista[j+1] = temp
			}
		}
	}
	fmt.Println("Lista de numeros ordenados: ", lista)
}
