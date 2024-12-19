package main

const quantidadeRepeticoes = 5

func Repetir(caractere string, quantidade int) string {
	var repeticoes string
	for i := 0; i < quantidade; i++ {
		repeticoes = repeticoes + caractere
	}
	return repeticoes
}