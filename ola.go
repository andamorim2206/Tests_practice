package main

import "fmt"


func Ola(name string) string {
	return "Olá, " + name
}

func main() {
	fmt.Println(Ola("chris"))
}