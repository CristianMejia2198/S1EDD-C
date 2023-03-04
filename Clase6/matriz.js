class nodoMatriz{
    constructor(posX, posY, posicion){
        this.siguiente = null;
        this.anterior = null;
        this.abajo = null;
        this.arriba = null;
        this.posX = posX;
        this.posY = posY;
        this.posicion = posicion;
    }
}

class Matriz{
    constructor(){
        this.principal = new nodoMatriz(-1,-1,"Raiz")
    }

    buscarF(y){
        let aux = this.principal
        while(aux){
            if(aux.posY === y){
                return aux;
            }else{
                aux = aux.abajo;
            }
        }
        return null;
    }

    buscarC(x){
        let aux = this.principal;
        while(aux){
            if(aux.posX === x){
                return aux;
            }else{
                aux = aux.siguiente
            }
        }
        return null;
    }

    insertarColumna(posicion,texto){
        const nuevoNodo = new nodoMatriz(posicion,-1,texto);
        let piv = this.principal;
        let pivA = this.principal;
        while(piv.siguiente){
            if(nuevoNodo.posX > piv.posX){
                pivA = piv;
                piv = piv.siguiente
            }else{
                nuevoNodo.siguiente = piv;
                nuevoNodo.anterior = pivA;
                pivA.siguiente = nuevoNodo;
                piv.anterior = nuevoNodo;
                return;
            }
        }
        nuevoNodo.anterior = piv;
        piv.siguiente = nuevoNodo;
    }

    insertarFila(posicion,texto){
        const nuevoNodo = new nodoMatriz(-1,posicion,texto);
        let piv = this.principal;
        let pivA = this.principal;
        while(piv.abajo){
            if(nuevoNodo.posY > piv.posY){
                pivA = piv;
                piv = piv.abajo;
            }else{
                nuevoNodo.abajo = piv;
                nuevoNodo.arriba = pivA;
                pivA.abajo = nuevoNodo;
                piv.arriba = nuevoNodo;
                return;
            }
        }
        nuevoNodo.arriba = piv;
        piv.abajo = nuevoNodo;
    }
    
    insertarNodo(x,y,texto){
        const nuevoNodo = new nodoMatriz(x,y,texto);
        let tempX = this.principal;
        let tempY = this.principal;
        //Agregar en Columna
        while(tempX.siguiente){
            if(tempX.posX === nuevoNodo.posX){
                break;
            }
            tempX = tempX.siguiente;
        }
        while(true){
            if(tempX.posY === nuevoNodo.posY){
                break;
            }else if(tempX.abajo !== null && tempX.abajo.posY > nuevoNodo.posY){
                nuevoNodo.abajo = tempX.abajo;
                nuevoNodo.arriba = tempX;
                tempX.abajo = nuevoNodo;
                break;
            }else if(tempX.abajo === null){
                nuevoNodo.arriba = tempX
                nuevoNodo.abajo = tempX.abajo
                tempX.abajo = nuevoNodo;
                break;
            }else{
                tempX = tempX.abajo;
            }
        }
        //Agregar en Fila
        while(tempY.abajo){
            if(tempY.posY === nuevoNodo.posY){
                break;
            }
            tempY = tempY.abajo;
        }
        while(true){
            if(tempY.posX === nuevoNodo.posX){
                break;
            }else if(tempY.siguiente !== null && tempY.siguiente.posX > nuevoNodo.posX){
                nuevoNodo.siguiente = tempY.siguiente;
                nuevoNodo.anterior = tempY;
                tempY.siguiente = nuevoNodo;
            }else if(tempY.siguiente === null){
                nuevoNodo.anterior = tempY;
                nuevoNodo.siguiente = tempY.siguiente;
                tempY.siguiente = nuevoNodo;
            }else{
                tempY = tempY.siguiente;
            }
        }
    }

    insertarElemento(x,y){
        let texto = x + "," + y;
        let nuevaFila = this.buscarF(y);
        let nuevaColumna = this.buscarC(x);
        /** Fila y Columna no existen */
        if(nuevaFila === null && nuevaColumna === null){
            this.insertarColumna(x, "C"+x);
            this.insertarFila(y, "F"+y);
            this.insertarNodo(x,y,texto);
        }else if(nuevaFila === null && nuevaColumna !== null){ /* Fila no existe, Columna si existe */
            this.insertarFila(y,"F"+y);
            this.insertarNodo(x,y,texto);
        }else if(nuevaFila !== null && nuevaColumna === null){/* Fila si existe, Columna no existe */
            this.insertarColumna(x, "C"+x);
            this.insertarNodo(x,y,texto);
        }else if(nuevaFila !== null && nuevaColumna !== null){/* Fila si existe, Columna si existe */
            this.insertarNodo(x,y,texto);
        }else{
            console.log("Me dio Ansiedad :(");
        }
    }

    reporte(){
        let cadena = "";
        let aux1 = this.principal;
        let aux2 = this.principal;
        let aux3 = this.principal;
        if(aux1 !== null){
            cadena = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n";
            /** Creacion de los nodos actuales */
            while(aux1){
                cadena += "nodo" + (aux1.posX+1) + (aux1.posY+1) + "[label=\"" + aux1.posicion + "\" ,rankdir=LR,group=" + (aux1.posX+1) + "]; \n";
                aux1 = aux1.siguiente;
            }
            cadena += "}"
            while(aux2){
                aux1 = aux2;
                cadena += "{rank=same; \n";
                while(aux1){
                    cadena += "nodo" + (aux1.posX+1) + (aux1.posY+1) + "[label=\"" + aux1.posicion + "\" ,group=" + (aux1.posX+1) + "]; \n";
                    aux1 = aux1.siguiente;
                }
                cadena += "}";
                aux2 = aux2.abajo;
            }
            /** Conexiones entre los nodos de la matriz */
            aux2 = aux3;
            while(aux2){
                aux1 = aux2;
                while(aux1.siguiente){
                    cadena += "nodo" + (aux1.posX+1) + (aux1.posY+1) + " -> " + "nodo" + (aux1.siguiente.posX+1) + (aux1.siguiente.posY+1) + " [dir=both];\n"
                    aux1 = aux1.siguiente
                }
                aux2 = aux2.abajo;
            }
            aux2 = aux3;
            while(aux2){
                aux1 = aux2;
                while(aux1.abajo){
                    cadena += "nodo" + (aux1.posX+1) + (aux1.posY+1) + " -> " + "nodo" + (aux1.abajo.posX+1) + (aux1.abajo.posY+1) + " [dir=both];\n"
                    aux1 = aux1.abajo
                }
                aux2 = aux2.siguiente;
            }
            cadena +=  "}";
        }else{
            cadena = "No hay elementos en la matriz"
        }
        return cadena;
    }
}

const matriz = new Matriz();
/* Elementos no existen */
matriz.insertarElemento(4,4);
matriz.insertarElemento(6,6);
matriz.insertarElemento(8,8);
matriz.insertarElemento(5,5);
matriz.insertarElemento(10,10);
/* fila y columna existen*/
matriz.insertarElemento(10,8);
matriz.insertarElemento(8,5);
/** fila si, columna no */
matriz.insertarElemento(1,8);
matriz.insertarElemento(3,8);
/** fila no, columna si */
matriz.insertarElemento(4,1);
matriz.insertarElemento(5,11);
document.getElementById("ejemplo").innerHTML = matriz.reporte();
console.log(matriz.reporte());

/**
<details>
<summary> Descripcion 🧐</summary>
<br>

</details> 
 */