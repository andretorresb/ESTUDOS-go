package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello string = "Hello, World!"
	var question string = "How are you today?"

	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))
	fmt.Println(strings.Contains(meet, "World"))
}
