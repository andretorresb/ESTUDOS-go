package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Quando é sábado?")
	today := time.Now().Weekday()

	//switch expressao

	switch time.Saturday {
	case today + 0:
		fmt.Println("Hoje é sábado!")
	case today + 1:
		fmt.Println("é amanhã!")
	case today + 2:
		fmt.Println("é em dois dias!")
	default:
		fmt.Println("Hoje não é sábado.")
	}
}
