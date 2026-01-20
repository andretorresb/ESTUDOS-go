package main

import "fmt"

func main() {
	resultado := soma(3, 5)
	fmt.Println("Resultado da soma:", resultado)

}

func soma(a int, b int) int {
	return a + b
}
