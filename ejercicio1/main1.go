package main

import "fmt"

func EsPalindromo(palabra string) bool{

	letras := []byte(palabra) 
    var invertida []byte 

    for i  := len(letras) - 1; i >= 0; i -- {
	invertida = append(invertida,  letras[i])
    }

	return string(palabra) == string(invertida)

}

func main()  {

   var palabra string 

   fmt.Println("Escribe una palabra: ") 
   fmt.Scanln(&palabra) 

   if EsPalindromo(palabra){
   fmt.Println("La palabra es un Palindromo.")
   }else{
	fmt.Println("La palabra No es un Palindromo.")
   }

 
}