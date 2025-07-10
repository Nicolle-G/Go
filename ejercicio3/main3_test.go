package main

import (
	"testing"
)

func TestSumar(t *testing.T) {

	resultado := sumar(10, 5)
	esperado := 15.0

	if resultado != esperado {
		t.Errorf("Suma incorrecta. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}

func TestRestar(t *testing.T) {

	resultado := restar(10, 5)
	esperado := 5.0

	if resultado != esperado {
		t.Errorf("Resta incorrecta. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}

func TestMultiplicar(t *testing.T) {

	resultado := multiplicar(10, 5)
	esperado := 50.0

	if resultado != esperado {
		t.Errorf("Multiplicacion incorrecta. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}

func TestDividir(t *testing.T) {

	resultado := dividir(10, 5)
	esperado := 2.0

	if resultado != esperado {
		t.Errorf("Division incorrecta. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}

func TestDividirPorCero(t *testing.T) {

	resultado := dividir(10, 0)
	esperado := 0.0

	if resultado != esperado {
		t.Errorf("Division por cero incorrecta. Esperado: %v , Obtenido: %v ", esperado, resultado)
	}
}

func TestTodasLasOperaciones(t *testing.T) {
	num1 := 10.0
	num2 := 5.0

	sumaEsperada := 15.0
	restaEsperada := 5.0
	multiplicacionEsperada := 50.0
	divisionEsperada := 2.0

	if resultado := sumar(num1, num2); resultado != sumaEsperada {
		t.Errorf("Suma incorrecta. Esperado: %v , Obtenido: %v ", sumaEsperada, resultado)
	}

	if resultado := restar(num1, num2); resultado != restaEsperada {
		t.Errorf("Resta incorrecta. Esperado: %v , Obtenido: %v ", restaEsperada, resultado)
	}

	if resultado := multiplicar(num1, num2); resultado != multiplicacionEsperada {
		t.Errorf("Multiplicacion incorrecta. Esperado: %v , Obtenido: %v ", multiplicacionEsperada, resultado)
	}

	if resultado := dividir(num1, num2); resultado != divisionEsperada {
		t.Errorf("Division incorrecta. Esperado: %v , Obtenido: %v ", divisionEsperada, resultado)
	}
}
