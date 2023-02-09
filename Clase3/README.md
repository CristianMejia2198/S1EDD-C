# Ejemplo Clase 3

## Ejemplo de Lista Doblemente Enlazada en Go

En este ejemplo se creo un programa el cual nos ayudara a guardar empleados segun su edad, usando una lista doble enlanzada para los trabajadores con edad menor a 25 años y una lista circular doblemente enlazada para los mayores de 25 años  
El codigo esta Separado de la siguiente manera  
__estructuras/empleado.go__ Esta la estructura del objeto que vamos a almacenar con sus atributos  
__estructuras/nodo.go__ Esta la estructura del objeto nodo de la lista  
__estructuras/interfaces.go__ Esta la estructura de la interfaz y asi mismo la declaracion de los metodos que utilizaremos  
__estructuras/listas.go__ Aqui se encuentra todo lo relacionado a listas, utilizacion de nodos, y funciones, etc    

### Nota Adicional
Si le da conflicto la ejecucion del programa en sus computadoras  
Borrar el archivo __go.mod__ y ejecutar en consola estando dentro de la carpeta Clase2 lo siguiente

```shell
    go mod init Clase3
```
Esto para crear el entorno dentro de sus computadoras y crear el archivo go.mod de nuevo  
Tener en cuenta el que archivo __go.mod__ debe quedar en la misma altura que el __main.go__

En consola dentro de la misma carpeta del archivo main.go ejecutar lo siguiente para ejecutar el programa
```shell
    go run main.go    
```