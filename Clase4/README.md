# Ejemplo Clase 4

## Ejemplo de Pila y Cola en Go

En este ejemplo se creo una pila y una cola que ambos guardaran una hora, esta hora se tomara desde la computadora, ademas de crear un menu principal y sub-menus para manejar esas estructuras.  
El codigo esta Separado de la siguiente manera  
__estructuras/cola.go__ Aqui se encuentra todo lo relacionado a la cola, utilizacion de nodos, y funciones, etc  
__estructuras/pila.go__ Aqui se encuentra todo lo relacionado a la pila, utilizacion de nodos, y funciones, etc  
__estructuras/nodo.go__ Esta la estructura del objeto nodo de la lista  
__estructuras/reportesGrafica.go__ En este archivo se encuentra las funciones necesarias para la creacion de archivos, los cuales nos ayudara para crear los archivos dot y asi mismo ejecutar el mismo para crear la imagen de las estructuras   

### Nota Adicional
Si le da conflicto la ejecucion del programa en sus computadoras  
Borrar el archivo __go.mod__ y ejecutar en consola estando dentro de la carpeta Clase2 lo siguiente

```shell
    go mod init Clase4
```
Esto para crear el entorno dentro de sus computadoras y crear el archivo go.mod de nuevo  
Tener en cuenta el que archivo __go.mod__ debe quedar en la misma altura que el __main.go__

En consola dentro de la misma carpeta del archivo main.go ejecutar lo siguiente para ejecutar el programa
```shell
    go run main.go    
```