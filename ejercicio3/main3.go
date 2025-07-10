package main

import "fmt"

func sumar(a, b float64) float64 {
	return a + b
}

func restar(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Por 0 no se puede divir")
		return 0
	}
	return a / b
}

func main() {

	var num1, num2 float64
	var operacion string
	var resultado float64

	fmt.Println("Calculadora")

	fmt.Println("Por favor ingrese el primer numero")
	fmt.Scanln(&num1)

	fmt.Println("Por favor ingrese el segundo numero")
	fmt.Scanln(&num2)

	fmt.Println("Que operacion deseas realizar? (+, -, *, /) o todas:")
	fmt.Scanln(&operacion)

	switch operacion {
	case "+":
		resultado = sumar(num1, num2)
		fmt.Println("El resultado de la suma es:", resultado)
	case "-":
		resultado = restar(num1, num2)
		fmt.Println("El resultado de la resta es:", resultado)
	case "*":
		resultado = multiplicar(num1, num2)
		fmt.Println("El resultado de la multiplicacion es:", resultado)
	case "/":
		resultado = dividir(num1, num2)
		fmt.Println("El resultado de la division es:", resultado)
	case "todas":
		fmt.Println("sumar: ", sumar(num1, num2))
		fmt.Println("restar ", restar(num1, num2))
		fmt.Println("multiplicar: ", multiplicar(num1, num2))
		fmt.Println("dividir: ", dividir(num1, num2))
	default:
		fmt.Println("Operacion no valida")
		return
	}

}
