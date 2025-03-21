package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 14)

	if total != 30 {
		t.Errorf("Resultado da soma invalido: %d | esperado: %d", total, 30)
	}
}
