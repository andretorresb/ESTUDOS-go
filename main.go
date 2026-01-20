package main

import "fmt"

type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	Estado string
	Cep    string
}

func main() {
	cliente1 := Cliente{
		Nome:  "Ana",
		Idade: 28,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 123,
			Estado: "SP",
			Cep:    "12345-678",
		},
	}

	cliente2 := Cliente{
		Nome:  "Andre",
		Idade: 21,
	}
	cliente2.Email = "andrelindao@gmail.com"

	cliente1.Endereco.Numero = 321

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
