package estructuras

type Operaciones interface {
	estaVacia() bool
	AgregarEmpleado(nombre string, edad int)
	newNodo(empleado *Empleado) *Nodo
	MostrarLista()
}
