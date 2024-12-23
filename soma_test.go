package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado '%d', esperado '%d', dado '%d'", resultado, esperado, numeros)
		}
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado '%d', esperado '%d', dado '%d'", resultado, esperado, numeros)
		}
	})
}

func TestSomaTodoOResto(t *testing.T) {

    verificaSomas := func(t *testing.T, resultado, esperado []int) {
        t.Helper()
        if !reflect.DeepEqual(resultado, esperado) {
            t.Errorf("resultado %v, esperado %v", resultado, esperado)
        }
    }

    t.Run("faz a soma do resto", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
        esperado := []int{2, 9}
        verificaSomas(t, resultado, esperado)
    })

    t.Run("soma slices vazios de forma segura", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
        esperado := []int{0, 9}
        verificaSomas(t, resultado, esperado)
    })

}
