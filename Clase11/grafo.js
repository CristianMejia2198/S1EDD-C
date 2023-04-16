class nodoMatrizAdyacencia{
    constructor(valor){
        this.siguiente = null
        this.abajo = null
        this.valor = valor
    }
}

class grafoDirigido{
    constructor(){
        this.principal = null
    }

    insertarF(texto){
        const nuevoNodo = new nodoMatrizAdyacencia(texto)
        if(this.principal === null){
            this.principal = nuevoNodo
        }else{
            let aux = this.principal
            while(aux.abajo){
                if(aux.valor === nuevoNodo.valor){
                    return
                }
                aux = aux.abajo
            }
            aux.abajo = nuevoNodo
        }
    }

    insertarC(padre, hijo){
        const nuevoNodo = new nodoMatrizAdyacencia(hijo)
        if(this.principal !== null && this.principal.valor === padre){
            let aux = this.principal
            while(aux.siguiente){
                aux = aux.siguiente
            }
            aux.siguiente = nuevoNodo
        }else{
            this.insertarF(padre)
            let aux = this.principal
            while(aux){
                if(aux.valor === padre){
                    break;
                }
                aux = aux.abajo
            }
            if(aux !== null){
                while(aux.siguiente){
                    aux = aux.siguiente
                }
                aux.siguiente = nuevoNodo
            }
        }
    }

    insertarValores(padre, hijos){
        let cadena = hijos.split(',')
        for(let i = 0; i < cadena.length; i++){
            this.insertarC(padre,cadena[i])
        }
    }

    grafica(){
        let cadena = "digraph grafoDirigido{ rankdir=LR; node [shape = circle];"
        let auxPadre = this.principal
        let auxHijo = this.principal
        while(auxPadre){
            auxHijo = auxPadre.siguiente
            while(auxHijo){
                cadena += auxPadre.valor + " -> " + auxHijo.valor + " "
                auxHijo = auxHijo.siguiente
            }
            auxPadre = auxPadre.abajo
        }
        cadena += "}"
        return cadena
    }
}

const grafo =  new grafoDirigido()

function insertar(){
    let padre = document.getElementById("padre").value 
    let hijos = document.getElementById("hijos").value 
    grafo.insertarValores(padre,hijos)
    refrescarGrafo()
}

function refrescarGrafo(){
    let url = 'https://quickchart.io/graphviz?graph=';
    let body = grafo.grafica()
    $("#image").attr("src", url + body);
    document.getElementById("padre").value = ""
    document.getElementById("hijos").value = ""
}