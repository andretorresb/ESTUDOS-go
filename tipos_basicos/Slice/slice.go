package main

import "fmt"

func main() {
	var gavetas []string
	gavetas = append(gavetas, "Copos", "Pratos", "Talheres")

	gavetas = gavetas[:2]
	fmt.Println(gavetas)
}

// SLICE [X: X-1]
