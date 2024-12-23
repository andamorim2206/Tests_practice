package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
    if nome == "" {
        nome = "Mundo"
    }

    return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
    switch idioma {
    case "Frances":
        prefixo = prefixoOlaFrances
    case "Espanhol":
        prefixo = prefixoOlaEspanhol
    default:
        prefixo = prefixoOlaPortugues
    }
    return
}

func main() {
	fmt.Println(Ola("chris", ""))
}
