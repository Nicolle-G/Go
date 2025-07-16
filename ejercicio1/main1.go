package main

import "fmt"

func IsPalindrome(word string) bool {

	letters := []rune(word)
	var reversed []rune

	for i := len(letters) - 1; i >= 0; i-- {
		reversed = append(reversed, letters[i])
	}

	return string(word) == string(reversed)

}

func main() {

	var word string

	fmt.Println("Write a Word: ")
	fmt.Scanln(&word)

	if IsPalindrome(word) {
		fmt.Println("The word is a Palindrome")
	} else {
		fmt.Println("The word is Not a Palindrome.")
	}

}
