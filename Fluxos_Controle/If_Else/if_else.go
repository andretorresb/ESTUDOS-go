package main

import "fmt"

func main() {
	players := map[string]int{
		"Andre": 30,
	}

	if value, ok := players["Andre"]; ok {
		fmt.Println("pontos:", value, ok)
	}
}
