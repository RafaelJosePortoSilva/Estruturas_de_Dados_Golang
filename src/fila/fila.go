package main

import (
	"errors"
	"fmt"
)

type Node struct {
	valor   int
	proximo *Node
}

func newNode(value int) *Node {
	return &Node{valor: value, proximo: nil}
}

type Fila struct {
	primeiro *Node
	ultimo   *Node
}

func newFila() Fila {
	return Fila{nil, nil}
}

func (f *Fila) estaVazia() bool {
	return (f.primeiro == nil) && (f.ultimo == nil)
}

func (f *Fila) enfileirar(value int) {
	no := newNode(value)

	if f.estaVazia() {
		f.primeiro, f.ultimo = no, no
	} else {
		ult := f.ultimo
		ult.proximo = no
		f.ultimo = no
	}
}

func (f *Fila) desenfileirar() (valor int, erro error) {
	erro, valor = nil, 0

	if !f.estaVazia() {
		valor = f.primeiro.valor
		f.primeiro = f.primeiro.proximo

		if f.primeiro == nil {
			f.ultimo = nil
		}

	} else {
		erro = errors.New("A fila está vazia")
	}

	return
}

func main() {

	fila := newFila()

	fila.enfileirar(1)
	fila.enfileirar(2)
	fila.enfileirar(3)
	fila.enfileirar(4)
	fila.enfileirar(5)
	fila.enfileirar(6)

	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())
	fmt.Println(fila.desenfileirar())

}
