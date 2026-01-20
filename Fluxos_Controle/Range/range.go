package main

import "fmt"

func main() {
	users := map[string]string{
		"nome":  "Andre",
		"idade": "21",
	}
	// _ é usado para ignorar o valor e pegar apenas a chave ou vice-versa
	for key, _ := range users {
		fmt.Println(key)
	}
}
