# Ejemplo Clase 5

## Ejemplo de Lista de Pilas, Lectura de CSV y Ordenamiento de Burbuja

En este ejemplo se creo una pila y una cola que ambos guardaran una hora, esta hora se tomara desde la computadora, ademas de crear un menu principal y sub-menus para manejar esas estructuras.  
El codigo esta Separado de la siguiente manera  
__clase/Lista.go__ Aqui se encuentra todo lo relacionado a la lista, utilizacion de nodos, y funciones, etc  
__clase/Pila.go__ Aqui se encuentra todo lo relacionado a la pila, utilizacion de nodos, y funciones, etc  
__clase/nodo.go__ Esta la estructura del objeto nodo de la lista y pila  
__clase/ordenamiento.go__ En este archivo se encuentra las funciones para el ordenamiento de burbuja para vectores y para la lista dinamica  
__clase/leerCSV.go__ Es este archivo se encuentra las funciones que se necesitan para la lectura, captura de datos de un archivo CSV, adicional a eso creacion de un ejemplo de archivo con extension JSON     

### Nota Adicional
Si le da conflicto la ejecucion del programa en sus computadoras  
Borrar el archivo __go.mod__ y ejecutar en consola estando dentro de la carpeta Clase5 lo siguiente

```shell
    go mod init Clase4
```
Esto para crear el entorno dentro de sus computadoras y crear el archivo go.mod de nuevo  
Tener en cuenta el que archivo __go.mod__ debe quedar en la misma altura que el __main.go__

En consola dentro de la misma carpeta del archivo main.go ejecutar lo siguiente para ejecutar el programa
```shell
    go run main.go    
```