package main

import "testing"

func TestEspalindromo(t *testing.T) {

	if EsPalindromo("oso") == false {
		t.Errorf("error oso deberia ser palindromo")

	}

	if EsPalindromo("casa") == true {
		t.Errorf("erro casa NO deberia ser palindromo")
	}

}
