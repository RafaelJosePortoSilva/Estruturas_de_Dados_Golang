package main

import (
	"errors"
	"fmt"
)

// Estrutura dos nós
type Node struct {
	valor   int
	proximo *Node
}

func newNode(value int) *Node {
	return &Node{value, nil}
}

// Estrutura e métodos da pilha
type Pilha struct {
	topo *Node
}

func newPilha() Pilha {
	return Pilha{nil}
}

func (p *Pilha) estaVazia() bool {
	return p.topo == nil
}

func (p *Pilha) empilhar(valor int) {

	no := newNode(valor)
	no.proximo = p.topo
	p.topo = no
}

func (p *Pilha) desempilhar() (valorNode int, erro error) {
	if !p.estaVazia() {
		valorNode = p.topo.valor
		erro = nil
		p.topo = p.topo.proximo
		return

	} else {
		valorNode = 0
		erro = errors.New("A pilha já está vazia!!")
		return
	}

}

func main() {
	var p Pilha

	p = newPilha()
	p.empilhar(5)
	fmt.Println(p.topo.valor)

	p.empilhar(7)
	fmt.Println(p.topo.valor)

	p.empilhar(9)
	fmt.Println(p.topo.valor)

	p.empilhar(11)
	fmt.Println(p.topo.valor)

	r, e := p.desempilhar()
	fmt.Println(r, e)

	r, e = p.desempilhar()
	fmt.Println(r, e)

	r, e = p.desempilhar()
	fmt.Println(r, e)

	r, e = p.desempilhar()
	fmt.Println(r, e)

	r, e = p.desempilhar()
	fmt.Println(r, e)
  
}
