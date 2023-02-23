package clase

import "math/rand"

type Pila struct {
	primero  *nodoPila
	longitud int
}

func (p *Pila) AgregarValores(cantidad int) {
	for i := 0; i < cantidad; i++ {
		p.push(rand.Intn(101))
	}
}

func (p *Pila) push(numero int) {
	if p.longitud == 0 {
		nuevoNodo := &nodoPila{valor: numero, siguiente: nil}
		p.primero = nuevoNodo
		p.longitud++
	} else {
		nuevoNodo := &nodoPila{valor: numero, siguiente: p.primero}
		p.primero = nuevoNodo
		p.longitud++
	}
}
