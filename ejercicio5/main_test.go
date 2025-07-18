package main

import "testing"

func TestInsertSearchValue(t *testing.T) {

	InsertValue("mateo", 10)
	InsertValue("oño", 20)

	value := SearchValue("mateo")
	if value != 10 {
		t.Errorf("Expected 10, but got it %d", value)
	}

	value = SearchValue("oño")
	if value != 20 {
		t.Errorf("Expected 20, but got it %d", value)
	}
}

func TestSearchClueNotExist(t *testing.T) {

	InsertValue("lupe", 5)

	value := SearchValue("Not Exist")
	if value != -1 {
		t.Errorf("Expected -1 for non-existent clue, but got it %d", value)
	}
}

func TestCollision(t *testing.T) {

	InsertValue("lupe", 5)
	InsertValue("pelu", 11)

	value1 := SearchValue("lupe")
	if value1 != 5 {
		t.Errorf("Expected 5 for 'Lupe', but it was obtained %d", value1)
	}

	value2 := SearchValue("pelu")
	if value2 != 11 {
		t.Errorf("Expected 11 for 'pelu', but it was obtained %d", value2)
	}
}
