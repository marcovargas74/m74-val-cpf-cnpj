package main

import "fmt"

//INTERFACES----
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type client struct {
	nome  string
	saldo float64
}

func (p *pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (c *client) toString() string {
	return fmt.Sprintf("cliente:%s - Saldo %.2f", c.nome, c.saldo)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

/*HOW TO USE INTERFACE*/
func main() {
	var coisa imprimivel = &pessoa{"Roberto", "Silva"}
	imprimir(coisa)
	coisa = &client{"Roberto", 1500.00}
	imprimir(coisa)

}
