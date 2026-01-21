package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Com o Ponteiro * antes do tipo Pessoa, o método pode modificar o objeto original.
func (p Pessoa) Apresentar() {
	p.Nome = "João" // Modificação local, não afeta o objeto original
	fmt.Printf("Olá, meu nome é %s, e eu tenho %d anos.", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Ana", Idade: 28}
	p1.Apresentar()
}
