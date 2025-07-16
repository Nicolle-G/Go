package main

import "testing"

func TestIspalindrome(t *testing.T) {

	if IsPalindrome("oño") == false {
		t.Errorf("bug oño should be palindrome")

	}

	if IsPalindrome("house") == true {
		t.Errorf("bug house NOT should be palindrome")
	}

}
