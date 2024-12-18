package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("chris")
	esperado := "Olá, chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}