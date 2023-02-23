package clase

type nodoLista struct {
	valor     int
	siguiente *nodoLista
	pila      *Pila
}

type nodoPila struct {
	valor     int
	siguiente *nodoPila
}
