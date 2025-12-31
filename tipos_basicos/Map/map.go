package main

import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas["Andre"] = 21
	pessoas["Yasmin"] = 19

	if idade, ok := pessoas["Yasmin"]; ok {
		fmt.Println("Essa pessoa existe no mapa", idade, ok)
	} else {
		fmt.Println("Essa pessoa nao existe no mapa")
	}
	delete(pessoas, "Andre")
	fmt.Println(pessoas)
}
