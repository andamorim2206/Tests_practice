package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func Ola(name string) string {
	if name == "" {
		name = "Mundo"
	}

	return prefixoOlaPortugues + name
}

func main() {
	fmt.Println(Ola("chris"))
}
