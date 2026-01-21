package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "Ana"}
	var p2 Pessoa = Pessoa{Nome: "Lais"}
	fmt.Println("nome antigo:", p1)

	var p3 *Pessoa = &p1
	p3.Nome = "Maria"
	fmt.Println(p1)
	fmt.Println(p2)
}
