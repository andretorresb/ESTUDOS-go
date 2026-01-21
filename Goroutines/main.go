package main

import (
	"fmt"
	"time"
)

func exibirMsg() {
	fmt.Println("Olá, mundo das Goroutines!")
}

func main() {
	go exibirMsg()
	time.Sleep(1 * time.Second)
	fmt.Println("Olá main function!")
}
