package main

import "testing"


func TestAdicionador(t *testing.T) {
	soma := Adicionar(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
	
}