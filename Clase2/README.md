# Ejemplo Clase 2

## Ejemplo de Lista Doblemente Enlazada en Go

En este ejemplo se realizo una lista doblemente enlazada en Lenguaje GO este consta de funciones basicas de insercion de nodo al final.  
El codigo esta Separado de la siguiente manera  
__Estructuras/alumno.go__ Esta la estructura de objeto que vamos a almacenar  
__Estructuras/nodo.go__ Esta la estructura del objeto nodo de la lista  
__Estructuras/ListaDoble.go__ Aqui se encuentra todo lo relacionado a lista, utilizacion de nodos, y funciones  

### Nota Adicional
Si le da conflicto la ejecucion del programa en sus computadoras  
Borrar el archivo __go.mod__ y ejecutar en consola estando dentro de la carpeta Clase2 lo siguiente

```shell
    go mod init Clase2
```
Esto para crear el entorno dentro de sus computadoras y crear el archivo go.mod de nuevo  
Tener en cuenta el que archivo __go.mod__ debe quedar en la misma altura que el __main.go__

En consola dentro de la misma carpeta del archivo main.go ejecutar lo siguiente  
```shell
    go run main.go    
```