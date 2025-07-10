package main

import "fmt"

func RingBuffer(datos []string, tamaño int) [5]string {
	var buffer [5]string
	writeIndex := 0

	for _, letra := range datos {
		buffer[writeIndex] = letra
		writeIndex = (writeIndex + 1) % tamaño
	}
	return buffer
}

func main() {
	var datos []string
	fmt.Println("Ingresa 10 letras una por una: ")

	for i := 0; i < 10; i++ {
		var letra string
		fmt.Println("Letra: ", i+1)
		fmt.Scanln(&letra)
		datos = append(datos, letra)
	}

	buffer := RingBuffer(datos, 5)

	fmt.Println("Contenido final del Ring Buffer: ")
	for i, val := range buffer {
		fmt.Println(i, val)
	}

}
