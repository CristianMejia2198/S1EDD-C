package clase

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}
	defer file.Close()

	leer := csv.NewReader(file)
	leer.Comma = ','
	encabezado := true
	for {
		linea, err := leer.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer la linea")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		fmt.Println("Nombre: ", linea[1], " Carnet: ", linea[0])
	}
}

func ArchivoJSON(lista *ListaNumero) string {
	contenido := "{\n"
	contenido += "\t\"Numeros\": [\n"
	aux := lista.Inicio
	for aux.siguiente != nil {
		contenido += "\t\t{\n"
		contenido += "\t\t\t\"Valor\": " + strconv.Itoa(aux.valor) + ", \n"
		contenido += "\t\t\t\"Carpeta_Raiz\": \"/\" \n"
		contenido += "\t\t},\n"
		aux = aux.siguiente
	}
	//esto es para el ultimo elemento
	contenido += "\t\t{\n"
	contenido += "\t\t\t\"Valor\": " + strconv.Itoa(aux.valor) + ", \n"
	contenido += "\t\t\t\"Carpeta_Raiz\": \"/\" \n"
	contenido += "\t\t}\n"
	contenido += "\t]\n"
	contenido += "}"
	return contenido
}
