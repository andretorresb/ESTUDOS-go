package main

import "fmt"

func main() {
	var fixo = 5

	multiplica := func(x int) int {
		return x * fixo
	}

	resultado := multiplica(10)
	fmt.Println(resultado)

}
