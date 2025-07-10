package main

import "testing"

func TestRingBuffer(t *testing.T) {
	datos := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	esperado := [5]string{"F", "G", "H", "I", "J"}

	resultado := RingBuffer(datos, 5)

	if resultado != esperado {
		t.Errorf("Resultado incorrecto. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}
