class nodoArbol {
    constructor(valor){
        this.izquierdo = null;
        this.derecho = null;
        this.valor = valor;
    }
}

class Arbol {
    constructor(){
        this.raiz = null;
    }
 
    insertarNodo(nodo, raiz){
        if(raiz === null){
            raiz = nodo;
        }else{
            if(raiz.valor === nodo.valor){
                raiz.valor = nodo.valor
            }else if(raiz.valor < nodo.valor){      // 10, null
                raiz.derecho = this.insertarNodo(nodo, raiz.derecho)
            }else{
                raiz.izquierdo = this.insertarNodo(nodo, raiz.izquierdo)
            }
        }
        return raiz;
    }

    insertarValor(valor){
        const nuevoNodo = new nodoArbol(valor);
        this.raiz = this.insertarNodo(nuevoNodo, this.raiz);
    }

    recorridoInorden(raiz){
        var cadena = ""
        if(raiz !== null){
            cadena = cadena + "\""
            cadena = cadena + raiz.valor
            cadena = cadena + "\""
            if(raiz.izquierdo !== null){
                cadena = cadena + " -> "
                cadena = cadena + this.recorridoInorden(raiz.izquierdo)
            }
            if(raiz.derecho !== null){
                cadena = cadena + " -> "
                cadena = cadena + this.recorridoInorden(raiz.derecho)
            }
        }
        return cadena
    }

    recorridoPreOrden(raiz){
        var cadena = ""
        if(raiz !== null){
            if(raiz.izquierdo !== null){
                cadena += this.recorridoPreOrden(raiz.izquierdo)
                cadena += " -> "
            }
            cadena += "\""
            cadena += raiz.valor
            cadena += "\""
            if(raiz.derecho !== null){
                cadena += " -> "
                cadena += this.recorridoPreOrden(raiz.derecho)
            }
        }
        return cadena
    }

    recorridoPostOrden(raiz){
        var cadena = ""
        if(raiz !== null){
            if(raiz.izquierdo !== null){
                cadena += this.recorridoPostOrden(raiz.izquierdo)
                cadena += " -> "
            }
            if(raiz.derecho !== null){
                cadena += this.recorridoPostOrden(raiz.derecho)
                cadena += " -> "
            }
            cadena += "\""
            cadena += raiz.valor
            cadena += "\""
        }
        return cadena
    }

    recorridoArbol(){
        console.log(this.recorridoInorden(this.raiz))
        console.log(this.recorridoPreOrden(this.raiz))
        console.log(this.recorridoPostOrden(this.raiz))
    }

    retornarValoresArbol(raiz, id){
        var cadena = "";
        var numero = id + 1;
        if(raiz !== null){
            cadena += "\"";
            cadena += raiz.valor;
            cadena += "\" ;";
            if(!(raiz.izquierdo === null) && !(raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += this.retornarValoresArbol(raiz.izquierdo, numero)
                cadena += "\"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += this.retornarValoresArbol(raiz.derecho, numero)
                cadena += "{rank=same" + "\"" + raiz.izquierdo.valor + "\"" + " -> " + "\"" + raiz.derecho.valor + "\""  + " [style=invis]}; "
            }else if(!(raiz.izquierdo === null) && (raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += this.retornarValoresArbol(raiz.izquierdo, numero)
                cadena += "\"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += "x" + numero + "[style=invis]";
                cadena += "{rank=same" + "\"" + raiz.izquierdo.valor + "\"" + " -> " + "x" + numero + " [style=invis]}; "
            }else if((raiz.izquierdo === null) && !(raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += "x" + numero + "[style=invis]";
                cadena += "; \"";
                cadena += raiz.valor;
                cadena += "\" -> ";
                cadena += this.retornarValoresArbol(raiz.derecho, numero)
                cadena += "{rank=same" + " x" + numero + " -> \"" + raiz.derecho.valor + "\"" +  " [style=invis]}; "
            }
        }
        return cadena;
    }
    graficarArbol(){
        var cadena = ""
        if(this.raiz !== null){
            cadena += "digraph arbol {"
            cadena += this.retornarValoresArbol(this.raiz, 0)
            cadena += "}"
        }
        return cadena
    }

}


const arbolBinario = new Arbol();

function agregarVarios(){
    let valor = document.getElementById("valor").value
    let valores = valor.split(',')
    try{
        valores.forEach(element => {
            arbolBinario.insertarValor(element)
        });
    }catch(error){
        alert("Hubo un error")
    }
    refrescarArbol()
}

function agregarVariosNumeros(){
    let valor = document.getElementById("valor").value
    let valores = valor.split(',')
    try{
        valores.forEach(element => {
            arbolBinario.insertarValor(parseInt(element))
        });
    }catch(error){
        alert("Hubo un error")
    }
    refrescarArbol();
}

function refrescarArbol(){
    let url = 'https://quickchart.io/graphviz?graph=';
    let body = arbolBinario.graficarArbol();
    $("#image").attr("src", url + body);
    document.getElementById("valor").value = "";
}